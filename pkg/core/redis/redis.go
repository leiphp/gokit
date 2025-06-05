package redis

import (
	"context"
	redis "github.com/redis/go-redis/v9"
)

var Client *redis.Client

func Init(addr, password string, db int) {
	Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}

func Get(ctx context.Context, key string) (string, error) {
	return Client.Get(ctx, key).Result()
}

func Set(ctx context.Context, key string, value interface{}) error {
	return Client.Set(ctx, key, value, 0).Err()
}
