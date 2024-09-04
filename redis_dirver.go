package redisdriver

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisClient implements the RedisDriver interface using go-redis client.
type RedisClient struct {
	client *redis.Client
}

// NewRedisClient creates a new RedisClient with the given Redis options.
func NewRedisClient(addr, password string, db int) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	return &RedisClient{client: rdb}
}

// Set stores a key-value pair in Redis with an optional expiration time.
func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

// Get retrieves a value from Redis by its key.
func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // key does not exist
	}
	return val, err
}

// Del removes one or more keys from Redis.
func (r *RedisClient) Del(ctx context.Context, keys ...string) error {
	return r.client.Del(ctx, keys...).Err()
}
