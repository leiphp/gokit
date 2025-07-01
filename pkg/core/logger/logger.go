package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"time"
)

var Logger *zap.Logger

//func InitLogger() {
//	log, _ := zap.NewProduction()
//	Logger = log
//}

func InitLogger2() {
	// 确保 logs 目录存在
	_ = os.MkdirAll("logs", 0744)
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Encoding:    "json", // 可选："json" 或 "console"
		OutputPaths: []string{
			"stdout",       // 终端
			"logs/app.log", // 文件路径（确保 logs 目录存在）
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder, // 改成可读时间格式
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	log, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	Logger = log
}

func InitLogger() {
	// 日志文件路径格式，如 logs/20250701.log
	logPath := filepath.Join("logs", "%Y%m%d.log")

	// rotatelogs：按天切割，保留 7 天的日志
	writer, err := rotatelogs.New(
		logPath,
		rotatelogs.WithRotationTime(24*time.Hour), // 每天切割
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 最多保留 7 天
	)
	if err != nil {
		panic("初始化日志文件失败: " + err.Error())
	}

	// zapcore.WriteSyncer 封装 rotatelogs writer
	writeSyncer := zapcore.AddSync(writer)

	// 编码器配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// core: 写入文件 + 指定日志级别
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 也可以用 zapcore.NewConsoleEncoder
		writeSyncer,
		zapcore.InfoLevel,
	)

	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

func Info(msg string, fields ...zap.Field) {
	Logger.WithOptions(zap.AddCallerSkip(1)).Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Logger.WithOptions(zap.AddCallerSkip(1)).Error(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Logger.WithOptions(zap.AddCallerSkip(1)).Warn(msg, fields...)
}
