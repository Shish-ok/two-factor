package redis_dump

import (
	"time"
)

type RedisRecord struct {
	Key string        `json:"key"`
	Val string        `json:"val"`
	Ttl time.Duration `json:"ttl"`
}
