package configs

import (
	"github.com/spf13/viper"
	"log"
)

type Env string

func Load(env Env) *Config {
	var config = Config{}

	switch env {
	case Development:
		viper.SetConfigFile("config.dev.yaml")
	case Production:
		viper.SetConfigFile("config.prod.yaml")
	default:
		log.Fatal("Don`t support the file .config")
	}

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .config : ", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if config.Application.Env == "development" {
		log.Println("The App is running in development")
	}

	return &config
}
