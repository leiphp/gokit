package test

import (
	"context"
	"github.com/leiphp/gokit/pkg/core/logger"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	logger.InitLogger()
	defer logger.Logger.Sync() // 2. 保证程序退出前写入 buffer 内容

	logger.Logger.Info("服务启动")
	logger.Info("封装日志时间类型", zap.Time("time", time.Now()))
	logger.Info("封装日志字符串类型", zap.String("message", "hello"))

	ctx := context.WithValue(context.Background(), "trace_id", "abc-123-xyz")
	logger.Info("处理请求成功",
		zap.Int("code", 200),
		zap.String("trace_id", ctx.Value("trace_id").(string)),
	)
}

//常用 zap.Field 构造函数对照表
//数据类型	使用函数					示例
//string	zap.String(key, val)	zap.String("msg", "hello")
//int	    zap.Int(key, val)	    zap.Int("code", 200)
//int64  	zap.Int64(key, val)	    zap.Int64("ts", time.Now().Unix())
//bool	    zap.Bool(key, val)	    zap.Bool("ok", true)
//time.Time	zap.Time(key, val)	    zap.Time("now", time.Now())
//error	    zap.Error(err)	        zap.Error(err)
//float64	zap.Float64(key, val)	zap.Float64("cost", 3.14)
//duration	zap.Duration(key, val)	zap.Duration("elapsed", time.Since(start))
