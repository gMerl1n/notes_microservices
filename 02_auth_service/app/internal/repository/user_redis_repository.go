package repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type IRedisRepositoryUser interface {
	SaveUserByToken(ctx context.Context, RefreshToken, userUUID string, expiresAt time.Duration) error
	GetUserByToken(ctx context.Context, RefreshToken string) (string, error)
	RemoveUserByToken(ctx context.Context, RefreshToken string) error
}

type RedisRepositoryUser struct {
	client *redis.Client
}

func NewRedisStoreUser(client *redis.Client) *RedisRepositoryUser {
	return &RedisRepositoryUser{client: client}
}

func (r *RedisRepositoryUser) SaveUserByToken(ctx context.Context, refreshToken, userUUID string, expiresAt time.Duration) error {

	if err := r.client.HSet(ctx, refreshToken, "userUUID", userUUID, "expiresAt", expiresAt).Err(); err != nil {
		return err
	}

	timeExpireSession := time.Now().Local().Add(10 * time.Second)

	r.client.ExpireAt(ctx, refreshToken, timeExpireSession)

	return nil
}

func (r *RedisRepositoryUser) GetUserByToken(ctx context.Context, refreshToken string) (string, error) {

	sessByRToken, err := r.client.HGetAll(ctx, refreshToken).Result()
	if err != nil {
		return "", err
	}

	userUUID, ok := sessByRToken["userUUID"]
	if !ok {
		return "", err
	}

	return userUUID, nil
}

func (r *RedisRepositoryUser) RemoveUserByToken(ctx context.Context, refreshToken string) error {
	err := r.client.Del(ctx, refreshToken).Err()
	if err != nil {
		return err
	}

	return nil

}
