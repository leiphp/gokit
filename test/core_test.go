package test

import (
	"context"
	"github.com/leiphp/gokit/pkg/core/httpclient"
	"github.com/leiphp/gokit/pkg/core/mysql"
	"github.com/leiphp/gokit/pkg/core/redis"
	"testing"
)

func TestCore(t *testing.T) {
	redis.Init("localhost:6379", "", 0)
	mysql.Init("user:pwd@tcp(127.0.0.1:3306)/test")
	trace.Init("my-service")

	ctx, span := trace.StartSpan(context.Background(), "main-span")
	defer span.End()

	redis.Set(ctx, "key", "value")
	val, _ := redis.Get(ctx, "key")

	body, _ := httpclient.Get("https://httpbin.org/get")
}
