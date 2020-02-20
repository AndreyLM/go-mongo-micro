package logger

import (
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

// Init - init logger from main
func Init() {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	logger = zapLogger.Sugar()
}

// Get - gets logger
func Get() *zap.SugaredLogger {
	return logger
}

// Close - flushes any buffered log entries
func Close() {
	logger.Sync()
}
