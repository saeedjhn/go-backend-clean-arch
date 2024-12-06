package bootstrap

import (
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/adaptor/jsonfilelogger"
	"github.com/saeedjhn/go-backend-clean-arch/internal/contract"
)

func NewLogger(configApp configs.Application, configLogger jsonfilelogger.Config) contract.Logger {
	strategy := createEnvironmentStrategy(configApp.Env, configLogger)

	return jsonfilelogger.New(strategy).Configure()
}

func createEnvironmentStrategy(env configs.Env, config jsonfilelogger.Config) jsonfilelogger.EnvironmentStrategy {
	switch env {
	case configs.Production:
		return jsonfilelogger.NewProductionStrategy(config)
	default:
		return jsonfilelogger.NewDevelopmentStrategy(config)
	}
}

// func NewLogger(configApp configs.Application, configLogger jsonfilelogger.Config) contract.Logger {
// 	var strategy jsonfilelogger.EnvironmentStrategy
//
// 	strategy = jsonfilelogger.NewDevelopmentStrategy(configLogger)
// 	if configApp.Env == configs.Production {
// 		strategy = jsonfilelogger.NewProductionStrategy(configLogger)
// 	}
//
// 	return jsonfilelogger.New(strategy)
// }

// func NewLogger(configApp configs.Application, configLogger jsonfilelogger.Config) contract.Logger {
// 	strategies := map[configs.Env]func(jsonfilelogger.Config) jsonfilelogger.EnvironmentStrategy{
// 		configs.Production: func(config jsonfilelogger.Config) jsonfilelogger.EnvironmentStrategy {
// 			return jsonfilelogger.NewProductionStrategy(config)
// 		},
// 		configs.Development: func(config jsonfilelogger.Config) jsonfilelogger.EnvironmentStrategy {
// 			return jsonfilelogger.NewDevelopmentStrategy(config)
// 		},
// 	}
//
// 	createStrategy, exists := strategies[configApp.Env]
// 	if !exists {
// 		createStrategy = func(config jsonfilelogger.Config) jsonfilelogger.EnvironmentStrategy {
// 			return jsonfilelogger.NewDevelopmentStrategy(config)
// 		}
// 	}
//
// 	strategy := createStrategy(configLogger)
//
// 	return jsonfilelogger.New(strategy)
// }
