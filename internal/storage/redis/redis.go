package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func NewRedisStorage(addr string, db int) (*RedisStorage, error) {
	testClient := &RedisStorage{
		redis.NewClient(&redis.Options{
			Addr: addr,
			DB:   db,
		}),
	}

	_, err := testClient.Ping(context.Background()).Result()
	return testClient, err
}

type RedisStorage struct {
	*redis.Client
}
