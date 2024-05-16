package configs

import (
	"github.com/spf13/viper"
	"log"
)

type Env string

const (
	Development Env = "development"
	Production  Env = "production"
)

func Load(env Env) *Config {
	cfg := Config{}

	switch env {
	case Development:
		viper.SetConfigFile("development.yaml")
	case Production:
		viper.SetConfigFile("production.yaml")
	default:
		log.Fatal("Don`t support the file .cfg")
	}

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .cfg : ", err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if cfg.Application.Env == "development" {
		log.Println("The App is running in development cfg")
	}

	return &cfg
}
