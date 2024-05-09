package redis_client

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type ConfigRedis struct {
	addrRedis     string
	passwordRedis string
	dbRedis       int
}

func NewRedisConfig(addrRedis string, passwordRedis string, dbRedis string) (*ConfigRedis, error) {

	dbr, err := strconv.Atoi(dbRedis)

	if err != nil {
		return nil, err
	}

	return &ConfigRedis{
		addrRedis:     addrRedis,
		passwordRedis: passwordRedis,
		dbRedis:       dbr,
	}, nil
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
