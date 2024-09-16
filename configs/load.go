package configs

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Env string

func Load(env Env) (*Config, error) {
	var config = Config{}

	switch env {
	case Development:
		viper.SetConfigFile("config.dev.yaml")
	case Production:
		viper.SetConfigFile("config.prod.yaml")
	default:
		return &config, errors.New("don`t support the file .config")
	}

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return &config, fmt.Errorf("can't find the file .config : %w", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return &config, fmt.Errorf("environment can't be loaded: %w", err)
	}

	if config.Application.Env == "development" {
		log.Println("The App is running in development")
	}

	return &config, nil
}
