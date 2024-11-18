package configs

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type Option struct {
	Prefix      string
	Delimiter   string
	Separator   string
	FilePath    []string
	CallbackEnv func(string) string
}

func Load(option Option) (*Config, error) {
	var config = Config{}

	if len(option.FilePath) == 0 {
		return &config, errors.New("no configuration file paths provided")
	}

	for _, path := range option.FilePath {
		viper.SetConfigFile(path)
		if err := viper.MergeInConfig(); err != nil {
			return &config, fmt.Errorf("failed to load config file at '%s': %w", path, err)
		}
	}

	viper.AutomaticEnv()
	err := viper.Unmarshal(&config)
	if err != nil {
		return &config, fmt.Errorf("failed to unmarshal configuration into struct: %w", err)
	}

	return &config, nil
}

//
// func Load(option Option) (*Config, error) {
//	var config = Config{}
//
//	viper.SetConfigFile(option.FilePath)
//
//	viper.AutomaticEnv()
//	err := viper.ReadInConfig()
//	if err != nil {
//		return &config, fmt.Errorf("can't find the file .config : %w", err)
//	}
//
//	err = viper.Unmarshal(&config)
//	if err != nil {
//		return &config, fmt.Errorf("environment can't be loaded: %w", err)
//	}
//
//	return &config, nil
// }
