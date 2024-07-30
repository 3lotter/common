package redis

import (
	"context"
	"github.com/bsm/redislock"
	"time"
)

type RedisLock struct {
	locker *redislock.Client
}

func NewRedisLock(client redislock.RedisClient) *RedisLock {
	locker := redislock.New(client)
	return &RedisLock{locker: locker}
}

func (rl *RedisLock) ObtainLock(lockKey string, ttl time.Duration) (*redislock.Lock, error) {
	ctx := context.Background()
	return rl.locker.Obtain(ctx, lockKey, ttl, nil)
}

func (rl *RedisLock) ReleaseLock(lock *redislock.Lock) error {
	ctx := context.Background()
	return lock.Release(ctx)
}

func (rl *RedisLock) RefreshLock(lock *redislock.Lock, ttl time.Duration) error {
	ctx := context.Background()
	return lock.Refresh(ctx, ttl, nil)
}
