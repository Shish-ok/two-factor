package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"two-factor-auth/internal/config"
)

func NewRedisStorage(cfg config.ServiceConfiguration) (*RedisStorage, error) {
	testClient := &RedisStorage{
		redis.NewClient(&redis.Options{
			Addr: cfg.RedisConfig.Addr,
			DB:   0,
		}),
	}

	_, err := testClient.Ping(context.Background()).Result()
	return testClient, err
}

type RedisStorage struct {
	*redis.Client
}
