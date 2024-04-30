package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type ConfigRedis struct {
	addrRedis     string
	passwordRedis string
	dbRedis       string
}

func NewRedisConfig(addrRedis string, passwordRedis string, dbRedis string) *ConfigRedis {
	return &ConfigRedis{
		addrRedis:     addrRedis,
		passwordRedis: passwordRedis,
		dbRedis:       dbRedis,
	}
}

func NewRedisClient(cfg *ConfigRedis) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.addrRedis,
		Password: cfg.passwordRedis,
		DB:       cfg.dbRedis,
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}
	return client, nil
}

// redis:
//   addr_redis: "redis:6379"
//   db_redis: "0"
