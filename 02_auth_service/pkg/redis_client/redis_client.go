package redis_client

import (
	"context"

	"github.com/gMerl1n/notes_microservices/internal/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(ctx context.Context, cfg *config.ConfigRedis) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.AddrRedis,
		Password: cfg.PasswordRedis,
		DB:       cfg.DBRedis,
	})

	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}
	return client, nil
}

// redis:
//   addr_redis: "redis:6379"
//   db_redis: "0"
