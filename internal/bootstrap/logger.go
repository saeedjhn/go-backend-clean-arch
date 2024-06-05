package bootstrap

import "go-backend-clean-arch/internal/infrastructure/logger"

func newLogger(config logger.Config) *logger.Logger {
	return logger.New(config)
}
