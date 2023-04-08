package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitialzieLogger() {
	logger, _ := zap.NewProduction()
	Logger = logger
}