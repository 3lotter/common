package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type MyRedis struct {
	*redis.Client
}

// SetupRedis 初始化Redis连接
func SetupRedis(redisURL string) (*MyRedis, error) {
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}
	redisClient := redis.NewClient(opt)
	myRedis := &MyRedis{Client: redisClient}
	pong := myRedis.Ping(context.Background())
	if pong.Err() != nil {
		return nil, err
	}
	return myRedis, nil
}
