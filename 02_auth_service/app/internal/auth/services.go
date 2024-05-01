package auth

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/iriskin77/notes_microservices/app/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateUser(ctx context.Context, dto CreateUserDTO) (userUUID string, err error)
	Login(ctx context.Context, loginUser LoginUserDTO) (*jwt.Tokens, error)
	RefreshTokens(ctx context.Context, refreshToken string) (jwt.Tokens, error)
}

// type StoreRedis interface {
// 	SetSession(ctx context.Context, SID string, sess Session, lifetime time.Duration) error
// }

type service struct {
	// создаем структуру, которая принимает репозиторий для работы с БД
	storage         Storage
	tokenManager    jwt.TokenManager
	redis           RedisStorage
	logger          *slog.Logger
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewService(
	repo *Repository,
	tokenManager jwt.TokenManager,
	redis *RedisRepo, log *slog.Logger,
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
		fmt.Println(err.Error())
	}

	user := NewUser(dto)

	s.logger.Debug("generate password hash")
	err = user.GeneratePasswordHash()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	userUUID, err = s.storage.CreateUser(ctx, user)

	if err != nil {
		fmt.Println(err.Error())
	}

	return userUUID, nil
}

func (s *service) Login(ctx context.Context, loginUser LoginUserDTO) (*jwt.Tokens, error) {

	lgUser := LoginUser(loginUser)

	user, err := s.storage.GetByEmail(ctx, lgUser.Email)

	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(lgUser.Password)); err != nil {
		return nil, err
	}

	tokens, err := s.createSession(ctx, user.UUID)
	if err != nil {
		fmt.Println(err)
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
		ExpiresAt: time.Now().Add(s.refreshTokenTTL),
	}

	err = s.redis.SetSession(ctx, res.RefreshToken, session)

	return res, err
}

func (s *service) RefreshTokens(ctx context.Context, refreshToken string) (jwt.Tokens, error) {

	sessByRToken, err := s.redis.GetSession(ctx, refreshToken)
	if err != nil {
		fmt.Println(" s.redis.GetSession(ctx, refreshToken)", err.Error())
	}

	s.redis.DeleteSession(ctx, refreshToken)

	newTokens, err := s.createSession(ctx, sessByRToken.UserUUID)

	if err != nil {
		fmt.Println(err.Error())
	}

	return newTokens, nil
}
