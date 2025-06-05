package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() {
	log, _ := zap.NewProduction()
	Logger = log
}
