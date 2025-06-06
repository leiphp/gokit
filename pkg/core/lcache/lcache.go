package lcache

import (
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	c    *cache.Cache
	once sync.Once

	defaultExpire = 5 * time.Minute
	defaultClean  = 10 * time.Minute
)

// initCache 会确保只初始化一次缓存
func initCache() {
	once.Do(func() {
		c = cache.New(defaultExpire, defaultClean)
	})
}

// Init 初始化缓存，设置默认过期时间和清理周期
func Init(defaultExpiration, cleanupInterval time.Duration) {
	c = cache.New(defaultExpiration, cleanupInterval)
}

func Set(key string, value interface{}, d time.Duration) {
	initCache()
	c.Set(key, value, d)
}

func Get(key string) (interface{}, bool) {
	initCache()
	return c.Get(key)
}

func Delete(key string) {
	initCache()
	c.Delete(key)
}

func Flush() {
	initCache()
	c.Flush()
}
