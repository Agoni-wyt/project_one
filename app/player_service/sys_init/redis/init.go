package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"one_ser.com/config"
)

var (
	Red *redis.Client
)

func Init(cfg *config.RedisConfig) (err error) {
	ctx := context.Background()
	Red = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	_, err = Red.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
