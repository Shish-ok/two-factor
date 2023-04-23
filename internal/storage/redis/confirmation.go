package redis

import (
	"context"
	"fmt"
	"time"
	"two-factor-auth/internal/models/confirmations"
)

const (
	attemptsPrefix = "attempts_"
)

func (r *RedisStorage) AddConfirmation(ctx context.Context, conf confirmations.Confirmation, ttlCode time.Duration) error {
	pipe := r.TxPipeline()

	pipe.Set(ctx, conf.RequestUID.ToStr(), conf.AuthCode.ToStr(), ttlCode)

	attemptsKey := fmt.Sprint(attemptsPrefix, conf.RequestUID)
	pipe.Set(ctx, attemptsKey, 0, ttlCode)

	_, err := pipe.Exec(ctx)
	return err
}

func (r *RedisStorage) incrementAttempt(ctx context.Context, requestUID confirmations.UID) error {
	key := fmt.Sprint(attemptsPrefix, requestUID)
	_, err := r.Incr(ctx, key).Result()

	return err
}

func (r *RedisStorage) GetAttempts(ctx context.Context, requestUID confirmations.UID) (string, error) {
	key := fmt.Sprint(attemptsPrefix, requestUID)
	attempts, err := r.Get(ctx, key).Result()

	return attempts, err
}

func (r *RedisStorage) GetCodeByRequestUID(ctx context.Context, requestUID confirmations.UID) (confirmations.Code, error) {
	code, err := r.Get(ctx, requestUID.ToStr()).Result()
	if err != nil {
		return "", err
	}

	err = r.incrementAttempt(ctx, requestUID)
	if err != nil {
		return "", err
	}

	return confirmations.Code(code), nil
}

func (r *RedisStorage) DeleteConfirmationByRequestUID(ctx context.Context, requestUID confirmations.UID) error {
	attemptsKey := fmt.Sprint(attemptsPrefix, requestUID)
	return r.Del(ctx, attemptsKey, requestUID.ToStr()).Err()
}
