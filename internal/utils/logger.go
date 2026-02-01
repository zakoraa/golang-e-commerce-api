package utils

import "go.uber.org/zap"

func NewLoger() *zap.Logger {
	logger, _ := zap.NewProduction()
	return logger
}