package log

import "go.uber.org/zap"

func GetLogger() *zap.Logger {
	if logger == nil {
		InitLogger("dev", "./logs/service-logs.log")
	}

	return logger
}
