package services

import (
	"context"
	"fmt"
	"time"

	"github.com/gMerl1n/notes_microservices/app/internal/repository"
	"github.com/gMerl1n/notes_microservices/app/pkg/jwt"
	"github.com/gMerl1n/notes_microservices/app/pkg/logging"
	"golang.org/x/crypto/bcrypt"
)

type IServiceUser interface {
	CreateUser(ctx context.Context, name, surname, email, password, repeatPassword string, age int) (userUUID string, err error)
	Login(ctx context.Context, email, password string) (*jwt.Tokens, error)
	RefreshTokens(ctx context.Context, refreshToken string) (*jwt.Tokens, error)
}

type ServiceUser struct {
	// создаем структуру, которая принимает репозиторий для работы с БД
	repo            repository.IRepositoryUser
	tokenManager    jwt.TokenManager
	redis           repository.IRedisRepositoryUser
	logger          *logging.Logger
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewServiceUser(
	repo repository.IRepositoryUser,
	tokenManager jwt.TokenManager,
	redis repository.IRedisRepositoryUser,
	log *logging.Logger,
	accessTokenTTL time.Duration,
	refreshTokenTTL time.Duration) *ServiceUser {

	return &ServiceUser{
		repo:            repo,
		tokenManager:    tokenManager,
		redis:           redis,
		logger:          log,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL}
}

func (s *ServiceUser) CreateUser(ctx context.Context, name, surname, email, password, repeatPassword string, age int) (userUUID string, err error) {
	s.logger.Debug("check password and repeat password")
	if password != repeatPassword {
		return "", fmt.Errorf("password does not match repeat password")
	}

	s.logger.Debug("generate password hash")

	hashedPassword, err := generatePasswordHash(password)
	if err != nil {
		s.logger.Error("failed to generate hashed pass: %w", err)
		return "", err
	}

	userUUID, err = s.repo.CreateUser(ctx, name, surname, email, hashedPassword, age)

	if err != nil {
		s.logger.Error("failed to create user: %w", err)
		return "", err
	}

	return userUUID, nil
}

func generatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password due to error %w", err)
	}
	return string(hash), nil
}

func (s *ServiceUser) Login(ctx context.Context, email, password string) (*jwt.Tokens, error) {

	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		s.logger.Error("Failed to user by email", err)
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		s.logger.Error("Failed to compare password and repeated password", err)
		return nil, err
	}

	tokens, err := s.createRedisSession(ctx, user.UUID)
	if err != nil {
		s.logger.Error("Failed to create session %w", err)
		return nil, err
	}

	return tokens, nil
}

func (s *ServiceUser) createRedisSession(ctx context.Context, userUUID string) (*jwt.Tokens, error) {

	var (
		tokens jwt.Tokens
		err    error
	)

	tokens.AccessToken, err = s.tokenManager.NewJWT(userUUID)
	if err != nil {
		return nil, err
	}

	tokens.RefreshToken, err = s.tokenManager.NewRefreshToken()
	if err != nil {
		return nil, err
	}

	ExpiresAt := time.Duration(s.refreshTokenTTL) * time.Minute

	if err := s.redis.SaveUserByToken(ctx, tokens.RefreshToken, userUUID, ExpiresAt); err != nil {
		return nil, err
	}

	return &tokens, err
}

func (s *ServiceUser) RefreshTokens(ctx context.Context, refreshToken string) (*jwt.Tokens, error) {

	userUUID, err := s.redis.GetUserByToken(ctx, refreshToken)
	if err != nil {
		s.logger.Error("failed to get session by refresh token %w", err)
		return nil, err
	}

	ok := s.redis.RemoveUserByToken(ctx, refreshToken)

	if ok != nil {
		s.logger.Warn("failed to delete session by refresh token %w", err)
	}

	newTokens, err := s.createRedisSession(ctx, userUUID)

	if err != nil {
		s.logger.Error("failed to create session with new refresh token %w", err)
		return nil, err
	}

	return newTokens, nil
}
