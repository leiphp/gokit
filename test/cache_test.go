package test

import (
	"github.com/leiphp/gokit/pkg/core/lcache"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	lcache.Init(5*time.Minute, 10*time.Minute)
	lcache.Set("foo", "bar", lcache.DefaultExpiration)
	val, found := lcache.Get("foo")
}
