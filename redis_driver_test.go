package redisdriver_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"your-module-path/redisdriver"
)

// MockRedisClient is a mock of the RedisDriver interface
type MockRedisClient struct {
	mock.Mock
}

func (m *MockRedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	args := m.Called(ctx, key, value, expiration)
	return args.Error(0)
}

func (m *MockRedisClient) Get(ctx context.Context, key string) (string, error) {
	args := m.Called(ctx, key)
	return args.String(0), args.Error(1)
}

func (m *MockRedisClient) Del(ctx context.Context, keys ...string) error {
	args := m.Called(ctx, keys)
	return args.Error(0)
}

func TestRedisDriver(t *testing.T) {
	mockRedis := new(MockRedisClient)
	mockRedis.On("Set", mock.Anything, "my_key", "my_value", mock.Anything).Return(nil)
	mockRedis.On("Get", mock.Anything, "my_key").Return("my_value", nil)
	mockRedis.On("Del", mock.Anything, "my_key").Return(nil)

	ctx := context.Background()

	// Test Set
	err := mockRedis.Set(ctx, "my_key", "my_value", 10*time.Second)
	assert.NoError(t, err)

	// Test Get
	val, err := mockRedis.Get(ctx, "my_key")
	assert.NoError(t, err)
	assert.Equal(t, "my_value", val)

	// Test Del
	err = mockRedis.Del(ctx, "my_key")
	assert.NoError(t, err)
}
