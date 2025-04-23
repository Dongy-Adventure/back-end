package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type IRedisClient interface {
	SetEx(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Exists(ctx context.Context, key string) *redis.IntCmd
}
