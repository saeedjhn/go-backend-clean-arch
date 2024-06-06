package bootstrap

import "go-backend-clean-arch/internal/infrastructure/logger"

func NewLogger(config logger.Config) *logger.Logger {
	return logger.New(config)
}
