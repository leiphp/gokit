package test

import (
	"github.com/leiphp/gokit/pkg/core/cron"
	"github.com/leiphp/gokit/pkg/core/logger"
	"testing"
)

func TestCron(t *testing.T) {
	logger.InitLogger()
	cron.Init()

	// 添加定时任务（每分钟执行一次）
	cron.Add("* * * * *", func() {
		logger.Logger.Info("每分钟任务执行")
	})
}
