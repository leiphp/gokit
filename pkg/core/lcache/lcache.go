package lcache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var c *cache.Cache

// Init 初始化缓存，设置默认过期时间和清理周期
func Init(defaultExpiration, cleanupInterval time.Duration) {
	c = cache.New(defaultExpiration, cleanupInterval)
}

func Set(key string, value interface{}, d time.Duration) {
	c.Set(key, value, d)
}

func Get(key string) (interface{}, bool) {
	return c.Get(key)
}

func Delete(key string) {
	c.Delete(key)
}

func Flush() {
	c.Flush()
}
