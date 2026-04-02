package main

import "go.uber.org/zap"

func main() {
	// zap.NewProduction()创建的 Logger 在记录日志时会自动记录调用函数的信息、打日志的时间等
	logger, _ := zap.NewProduction()

	/*
		func (log *Logger) Info(msg string, fields ...Field) {
			if ce := log.check(InfoLevel, msg); ce != nil {
				ce.Write(fields...)
			}
		}
	*/
	logger.Info("Success..", zap.String("statusCode", "200"), zap.String("url", "www.baidu.com"))
}
