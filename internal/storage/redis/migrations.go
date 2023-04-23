package redis

import (
	"context"
	"two-factor-auth/internal/models/redis_dump"
)

func (r *RedisStorage) PutRecords(ctx context.Context, records []redis_dump.RedisRecord) {
	for _, record := range records {
		r.Set(ctx, record.Key, record.Val, record.Ttl)
	}
}

func (r *RedisStorage) GetAllRecords(ctx context.Context) ([]redis_dump.RedisRecord, error) {
	keys, err := r.Keys(ctx, "*").Result()
	if err != nil {
		r.Save(ctx)
		return nil, err
	}

	records := make([]redis_dump.RedisRecord, len(keys))
	for i, key := range keys {
		value, err := r.Get(ctx, key).Result()
		if err != nil {
			continue
		}
		ttl, err := r.TTL(ctx, key).Result()
		if err != nil {
			continue
		}

		records[i] = redis_dump.RedisRecord{
			Key: key,
			Val: value,
			Ttl: ttl,
		}
	}

	r.FlushAll(ctx)

	return records, nil
}
