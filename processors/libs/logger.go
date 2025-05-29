package libs

import (
	"processors/configs"
	"sync"

	"go.uber.org/zap"
)

var logger *zap.SugaredLogger
var logerOnce sync.Once

func LoggerInstance() *zap.SugaredLogger {
	logerOnce.Do(func() {
		var zapLogger *zap.Logger

		c := configs.GlobalConfigInstance()

		if c.Debug() {
			zapLogger, _ = zap.NewDevelopment()
			defer zapLogger.Sync()
		} else {
			zapLogger, _ = zap.NewProduction()
			defer zapLogger.Sync()
		}

		logger = zapLogger.Sugar()
	})
	return logger
}
