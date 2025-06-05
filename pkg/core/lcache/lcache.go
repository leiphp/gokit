package lcache

import (
	"github.com/rogpeppe/go-internal/cache"
	"time"
)

var c *cache.Cache

func Init(defaultExpiration, cleanupInterval time.Duration) {
	c = cache.New(defaultExpiration, cleanupInterval)
}

func Set(key string, value interface{}, d time.Duration) {
	c.Set(key, value, d)
}

func Get(key string) (interface{}, bool) {
	return c.Get(key)
}
