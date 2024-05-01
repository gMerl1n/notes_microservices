package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Session struct {
	RefreshToken string
	ExpiresAt    time.Time
}

type RedisStorage interface {
	SetSession(ctx context.Context, SID string, sess Session, lifetime time.Duration) error
	GetSession(ctx context.Context, SID string) (string, error)
	DeleteSession(ctx context.Context, SID string) error
}

type RedisRepo struct {
	client *redis.Client
}

func NewRedisStore(client *redis.Client) *RedisRepo {
	return &RedisRepo{client: client}
}

func (r *RedisRepo) SetSession(ctx context.Context, SID string, sess Session, lifetime time.Duration) error {
	tokenBytes, err := json.Marshal(sess)
	if err != nil {
		fmt.Println(err.Error())
	}
	if err := r.client.Set(ctx, SID, tokenBytes, lifetime).Err(); err != nil {
		return err
	}
	return nil
}

func (r *RedisRepo) GetSession(ctx context.Context, SID string) (string, error) {
	val, err := r.client.Get(ctx, SID).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func (r *RedisRepo) DeleteSession(ctx context.Context, SID string) error {
	err := r.client.Del(ctx, SID).Err()
	if err != nil {
		return err
	}
	return nil
}
