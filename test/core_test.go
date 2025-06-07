package test

import (
	"context"
	"fmt"
	"github.com/leiphp/gokit/pkg/core/mysql"
	"github.com/leiphp/gokit/pkg/core/redis"
	"github.com/leiphp/gokit/pkg/core/tracing"
	"testing"
)

func TestCore(t *testing.T) {
	redis.Init("localhost:6379", "", 0)
	mysql.Init("user:pwd@tcp(127.0.0.1:3306)/test")

	tracing.Init("my-service")
	ctx, span := tracing.StartSpan(context.Background(), "main-span")
	defer span.End()

	redis.Set(ctx, "key", "value")
	val, _ := redis.Get(ctx, "key")
	fmt.Println("val:", val)
}
