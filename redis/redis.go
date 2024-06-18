package redis

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
)

const (
	solo    = "solo"
	cluster = "cluster"

	hostSplit = ","
)

type MyRedis struct {
	redis.UniversalClient
}

type Conf struct {
	Type        string
	Host        string
	Password    string
	MaxActive   int
	MaxIdle     int
	IdleTimeout int
}

// SetupRedis 初始化Redis连接
func SetupRedis(conf Conf) (*MyRedis, error) {
	opts := redis.UniversalOptions{
		Password: conf.Password,

		MaxActiveConns:  conf.MaxActive,
		MaxIdleConns:    conf.MaxIdle,
		ConnMaxIdleTime: time.Duration(conf.IdleTimeout) * time.Second,
	}

	switch conf.Type {
	case solo:
		opts.Addrs = []string{conf.Host}
	case cluster:
		opts.Addrs = strings.Split(conf.Host, hostSplit)
	default:
		return nil, fmt.Errorf("redis type [%s] not supported", conf.Type)
	}

	redisClient := redis.NewUniversalClient(&opts)
	if _, err := redisClient.Ping(context.Background()).Result(); err != nil {
		redisClient.Close()
		return nil, errors.Wrap(err, "redisClient.Ping fail")
	}

	return &MyRedis{redisClient}, nil
}
