package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/iriskin77/notes_microservices/app/pkg/jwt"
	"github.com/iriskin77/notes_microservices/app/pkg/logging"
)

type Service interface {
	CreateUser(ctx context.Context, dto CreateUserDTO) (userUUID string, err error)
	Login(ctx context.Context, loginUser LoginUserDTO) (*jwt.Tokens, error)
	RefreshTokens(ctx context.Context, refreshToken string) (*jwt.Tokens, error)
}

type service struct {
	// создаем структуру, которая принимает репозиторий для работы с БД
	storage         Storage
	tokenManager    jwt.TokenManager
	redis           RedisStorage
	logger          *logging.Logger
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewService(
	repo *Repository,
	tokenManager jwt.TokenManager,
	redis *RedisRepo, log *logging.Logger,
	accessTokenTTL time.Duration,
	refreshTokenTTL time.Duration) *service {

	return &service{
		storage:         repo,
		tokenManager:    tokenManager,
		redis:           redis,
		logger:          log,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL}
}

func (s *service) CreateUser(ctx context.Context, dto CreateUserDTO) (userUUID string, err error) {
	s.logger.Debug("check password and repeat password")
	if dto.Password != dto.RepeatPassword {
		return "", fmt.Errorf("password does not match repeat password")
	}

	user := NewUser(dto)

	s.logger.Debug("generate password hash")
	err = user.GeneratePasswordHash()
	if err != nil {
		s.logger.Error("failed to generate hashed pass: %w", err)
		return "", err
	}

	userUUID, err = s.storage.CreateUser(ctx, user)

	if err != nil {
		s.logger.Error("failed to create user: %w", err)
		return "", err
	}

	return userUUID, nil
}

func (s *service) Login(ctx context.Context, loginUser LoginUserDTO) (*jwt.Tokens, error) {

	lgUser := LoginUser(loginUser)

	user, err := s.storage.GetByEmail(ctx, lgUser.Email)
	if err != nil {
		s.logger.Error("Failed to user by email", err)
		return nil, err
	}

	// ph, _ := generatePasswordHash(lgUser.Password)

	// fmt.Println("Login", user)
	// fmt.Println("Login", user.Password)

	// if err := bcrypt.CompareHashAndPassword([]byte(ph), []byte(lgUser.Password)); err != nil {
	// 	s.logger.Error("Failed to compare password and repeated password", err)
	// 	return nil, err
	// }

	tokens, err := s.createSession(ctx, user.UUID)
	if err != nil {
		s.logger.Error("Failed to create session %w", err)
		return nil, err
	}

	return &tokens, nil
}

func (s *service) createSession(ctx context.Context, userUUID string) (jwt.Tokens, error) {

	var (
		res jwt.Tokens
		err error
	)

	res.AccessToken, err = s.tokenManager.NewJWT(userUUID)
	if err != nil {
		return res, err
	}

	res.RefreshToken, err = s.tokenManager.NewRefreshToken()
	if err != nil {
		return res, err
	}
	session := Session{
		UserUUID:  userUUID,
		ExpiresAt: time.Duration(s.refreshTokenTTL) * time.Minute,
	}

	err = s.redis.SetSession(ctx, res.RefreshToken, session)

	return res, err
}

func (s *service) RefreshTokens(ctx context.Context, refreshToken string) (*jwt.Tokens, error) {

	sessByRToken, err := s.redis.GetSession(ctx, refreshToken)
	if err != nil {
		s.logger.Error("failed to get session by refresh token %w", err)
		return nil, err
	}

	ok := s.redis.DeleteSession(ctx, refreshToken)

	if ok != nil {
		s.logger.Warn("failed to delete session by refresh token %w", err)
	}

	newTokens, err := s.createSession(ctx, sessByRToken.UserUUID)

	if err != nil {
		s.logger.Error("failed to create session with new refresh token %w", err)
		return nil, err
	}

	return &newTokens, nil
}
