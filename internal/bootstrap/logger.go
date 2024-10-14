package bootstrap

import "github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/logger"

func NewLogger(c logger.Config) *logger.Logger {
	return logger.New(c)
}
