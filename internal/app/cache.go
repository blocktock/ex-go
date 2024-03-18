package app

import (
	"context"
	"ex-go/internal/config"
	"github.com/go-redis/redis/v8"
)

func InitCache() (*redis.Client, func(), error) {

	cfg := config.C.Redis

	// Redis 连接配置
	options := &redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	}
	client := redis.NewClient(options)

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, nil, err
	}

	cleanFunc := func() {
		//todo ... 错误处理
		_ = client.Close()
	}

	return client, cleanFunc, nil
}
