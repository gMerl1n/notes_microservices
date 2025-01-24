package repository

import (
	"context"
	"time"

	"github.com/gMerl1n/notes_microservices/pkg/logging"
	"github.com/redis/go-redis/v9"
)

type IRedisRepositoryUser interface {
	SaveUserByToken(ctx context.Context, RefreshToken, userID string, expiresAt time.Duration) error
	GetUserByToken(ctx context.Context, RefreshToken string) (string, error)
	RemoveUserByToken(ctx context.Context, RefreshToken string) error
}

type RedisRepositoryUser struct {
	client *redis.Client
	logger *logging.Logger
}

func NewRedisStoreUser(client *redis.Client, logger *logging.Logger) *RedisRepositoryUser {
	return &RedisRepositoryUser{client: client, logger: logger}
}

func (r *RedisRepositoryUser) SaveUserByToken(ctx context.Context, refreshToken, userID string, expiresAt time.Duration) error {

	r.logger.Info("Refresh Token:", refreshToken)

	if err := r.client.HSet(ctx, refreshToken, "userID", userID, "expiresAt", expiresAt).Err(); err != nil {
		return err
	}

	timeExpireSession := time.Now().Local().Add(10 * time.Minute)

	r.client.ExpireAt(ctx, refreshToken, timeExpireSession)

	return nil
}

func (r *RedisRepositoryUser) GetUserByToken(ctx context.Context, refreshToken string) (string, error) {

	sessByRToken, err := r.client.HGetAll(ctx, refreshToken).Result()
	if err != nil {
		return "", err
	}

	r.logger.Info("User ID from refresh token: ", sessByRToken["userID"])

	userID, ok := sessByRToken["userID"]
	if !ok {
		return "", err
	}

	return userID, nil
}

func (r *RedisRepositoryUser) RemoveUserByToken(ctx context.Context, refreshToken string) error {
	err := r.client.Del(ctx, refreshToken).Err()
	if err != nil {
		return err
	}

	return nil

}
