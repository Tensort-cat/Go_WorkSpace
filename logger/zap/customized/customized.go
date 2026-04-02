package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 把 zap 做进一步的自定义配置，让日志不光能输出到控制台，
// 也能输出到文件，再把日志时间由时间戳格式，
// 换成更容易被人类看懂的DateTime时间格式

var logger *zap.Logger

func init() {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 设置日志记录中时间的格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 日志Encoder 还是JSONEncoder，把日志行格式化成JSON格式的
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	file, _ := os.OpenFile("./tmp/test.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	fileWriteSyncer := zapcore.AddSync(file)

	core := zapcore.NewTee(
		// 同时向控制台和文件写日志， 生产环境记得把控制台写入去掉，日志记录的基本是Debug 及以上，生产环境记得改成Info
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel), // 写控制台
		zapcore.NewCore(encoder, fileWriteSyncer, zapcore.DebugLevel),            // 写文件
	)

	// 生产环境版本
	// core := zapcore.NewTee(
	// 	zapcore.NewCore(encoder, fileWriteSyncer, zapcore.InfoLevel), // 写文件
	// )

	logger = zap.New(core)
}

func main() {
	logger.Info("fuck you", zap.String("code", "200"), zap.String("url", "www.zhihu.com"))
}
