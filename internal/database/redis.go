package database

import (
	"context"
	"fmt"

	"github.com/Dongy-s-Advanture/back-end/internal/config"
	"github.com/redis/go-redis/v9"
)

func InitRedis(conf *config.DbConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddr,
		Password: conf.RedisPassword,
		DB:       conf.RedisDB,
	})
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error connecting redis: %v", err))
	}
	return rdb, err
}
