package test

import (
	"github.com/leiphp/gokit/pkg/core/mq"
	"testing"
)

func TestMq(t *testing.T) {
	mq.Init("amqp://guest:guest@localhost:5672/")
	mq.Publish("my-queue", []byte("hello world"))
	mq.Consume("my-queue", func(body []byte) {
		println("received:", string(body))
	})
}
