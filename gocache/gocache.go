package gocache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type MyCache struct {
	*cache.Cache
}

// SetupCache 初始化缓存
func SetupCache(defaultExpireTime, cleanupInterval int) MyCache {
	c := cache.New(time.Duration(defaultExpireTime)*time.Second, time.Duration(cleanupInterval)*time.Second)
	myCache := MyCache{c}
	return myCache
}
