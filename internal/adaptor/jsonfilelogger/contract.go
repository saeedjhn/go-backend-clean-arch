package jsonfilelogger

import "go.uber.org/zap"

type EnvironmentStrategy interface {
	CreateLogger() *zap.Logger
}
