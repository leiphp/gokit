package test

import (
	"fmt"
	"github.com/leiphp/gokit/pkg/core/lcache"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	lcache.Init(5*time.Minute, 10*time.Minute)
	lcache.Set("name", "leixiaotian", time.Minute)
	val, ok := lcache.Get("name")
	if ok {
		fmt.Println("缓存命中：", val)
	}
}
