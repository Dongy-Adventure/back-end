package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type goRedisAdapter struct {
	client *redis.Client
}

func NewGoRedisAdapter(client *redis.Client) IRedisClient {
	return &goRedisAdapter{client: client}
}

func (a *goRedisAdapter) SetEx(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return a.client.SetEx(ctx, key, value, expiration)
}

func (a *goRedisAdapter) Exists(ctx context.Context, key string) *redis.IntCmd {
	return a.client.Exists(ctx, key)
}

// Implement other methods you need from the Redis client
