package redisdriver

import (
	"context"
	"time"
)

// RedisDriver defines the interface for interacting with Redis.
type RedisDriver interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, keys ...string) error
}
