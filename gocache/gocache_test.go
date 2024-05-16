package gocache

import (
	"testing"
	"time"

	"github.com/patrickmn/go-cache"
)

// TestSetupCache tests the SetupCache function to ensure it initializes the cache correctly.
func TestSetupCache(t *testing.T) {
	tests := []struct {
		name              string
		defaultExpireTime int
		cleanupInterval   int
		expireAfter       int  // time after which item should expire
		shouldFind        bool // whether item should still be found in cache
	}{
		{
			name:              "Valid expiration",
			defaultExpireTime: 1, // 1 second
			cleanupInterval:   1,
			expireAfter:       2, // check after 2 seconds
			shouldFind:        false,
		},
		{
			name:              "Zero expiration",
			defaultExpireTime: 0, // no default expiration
			cleanupInterval:   1,
			expireAfter:       1,    // check after 1 second
			shouldFind:        true, // item should still be there because defaultExpireTime is 0
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			myCache := SetupCache(tt.defaultExpireTime, tt.cleanupInterval)
			myCache.Set("key", "value", cache.DefaultExpiration)

			time.Sleep(time.Duration(tt.expireAfter) * time.Second)
			_, found := myCache.Get("key")

			if found != tt.shouldFind {
				t.Errorf("%s: expected item present: %v, got: %v", tt.name, tt.shouldFind, found)
			}
		})
	}
}
