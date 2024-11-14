package bootstrap

import (
	logger2 "github.com/saeedjhn/go-backend-clean-arch/pkg/logger"
)

func NewLogger(c logger2.Config) *logger2.Logger {
	return logger2.New(c)
}
