package cache

import (
	"context"

	"github.com/go-redis/redis/v9"
	"github.com/yoshi429/test/config"
)

type RedisContext struct {
	RedisClient *redis.Client
	ctx         context.Context
}

func New(c config.Configs) *RedisContext {
	rds := redis.NewClient(&redis.Options{
		Addr:     "test-redis:6379",
		Password: "", // no password sret
		DB:       0,  // use default DB
	})

	return &RedisContext{
		RedisClient: rds,
		ctx:         context.Background(),
	}
}

func (r RedisContext) SET(key, value string) error {
	err := r.RedisClient.Set(r.ctx, key, value, 0).Err()
	return err
}

func (r RedisContext) GET(key string) (string, error) {
	val, err := r.RedisClient.Get(r.ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, err
}

func (r RedisContext) IsNotExistKey(err error) bool {
	if err == redis.Nil {
		return true
	}
	return false
}
