package auth

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Session struct {
	UserUUID  string
	ExpiresAt time.Duration
}

type RedisStorage interface {
	SetSession(ctx context.Context, RefreshToken string, sess Session) error
	GetSession(ctx context.Context, RefreshToken string) (Session, error)
	DeleteSession(ctx context.Context, RefreshToken string) error
}

type RedisRepo struct {
	client *redis.Client
}

func NewRedisStore(client *redis.Client) *RedisRepo {
	return &RedisRepo{client: client}
}

func (r *RedisRepo) SetSession(ctx context.Context, RefreshToken string, sess Session) error {
	if err := r.client.HSet(ctx, RefreshToken, "UserUUID", sess.UserUUID, "ExpiresAt", sess.ExpiresAt).Err(); err != nil {
		return err
	}

	timeExpireSession := time.Now().Local().Add(10 * time.Second)

	r.client.ExpireAt(ctx, RefreshToken, timeExpireSession)

	return nil
}

func (r *RedisRepo) GetSession(ctx context.Context, RefreshToken string) (Session, error) {

	sess := Session{}

	sessByRToken, err := r.client.HGetAll(ctx, RefreshToken).Result()
	if err != nil {
		return sess, err
	}

	sess.UserUUID = sessByRToken["UserUUID"]

	return sess, nil
}

func (r *RedisRepo) DeleteSession(ctx context.Context, RefreshToken string) error {
	err := r.client.Del(ctx, RefreshToken).Err()
	if err != nil {
		return err
	}

	return nil

}
