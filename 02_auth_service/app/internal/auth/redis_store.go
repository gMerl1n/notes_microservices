package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Session struct {
	UserUUID  string
	ExpiresAt time.Time
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
	if err := r.client.HMSet(ctx, RefreshToken, "UserUUID", sess.UserUUID, "ExpiresAt", sess.ExpiresAt).Err(); err != nil {
		return err
	}
	return nil
}

func (r *RedisRepo) GetSession(ctx context.Context, RefreshToken string) (Session, error) {

	sess := Session{}

	sessByRToken, err := r.client.HGetAll(ctx, RefreshToken).Result()
	if err != nil {
		return sess, err
	}

	//fmt.Println(sessByRToken["ExpiresAt"])

	sess.UserUUID = sessByRToken["UserUUID"]
	// sess.ExpiresAt, err = time.Parse(time.RFC1123, sessByRToken["ExpiresAt"])
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	return sess, nil
}

func (r *RedisRepo) DeleteSession(ctx context.Context, RefreshToken string) error {
	fmt.Println("DeleteSession")
	err := r.client.Del(ctx, RefreshToken).Err()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
