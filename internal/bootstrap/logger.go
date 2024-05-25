package bootstrap

import "go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/logger"

func newLogger(config logger.Config) *logger.Logger {
	return logger.New(config)
}
