package test

import (
	"github.com/leiphp/gokit/pkg/core/logger"
	"testing"
)

func TestLogger(t *testing.T) {
	logger.InitLogger()
	logger.Logger.Info("服务启动")
}
