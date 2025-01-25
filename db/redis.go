package db

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"os"
	"time"
)

var ctx = context.Background()

func NewRedisClient() *redis.Client {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return rdb
}

func SetCache(rdb *redis.Client, key string, value string, ttl time.Duration) error {
	err := rdb.Set(ctx, key, value, ttl).Err()
	return err
}

func GetCache(rdb *redis.Client, key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", nil
	}
	return val, err
}
