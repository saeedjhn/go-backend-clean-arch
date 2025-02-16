package setuptest

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

var m sync.Mutex //nolint:gochecknoglobals // nothing

func LoadConfig[T any](path string) (T, error) {
	m.Lock()
	defer m.Unlock()

	var config T

	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return config, fmt.Errorf("can't find the file: %w", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("environment can't be loaded: %w", err)
	}

	return config, nil
}

// const _fileExtension = "yml"
// func LoadConfig(path string) (*configs.Config, error) {
// 	filesWithExt, err := configs.CollectFilesWithExt(
// 		path,
// 		_fileExtension,
// 	)
// 	if err != nil {
// 		return nil, fmt.Errorf(
// 			"unexpected error while loading configuration files from directory: %s. Error: %w",
// 			path,
// 			err,
// 		)
// 	}
//
// 	cfgOption := configs.Option{
// 		Prefix:      "",
// 		Delimiter:   "",
// 		Separator:   "",
// 		FilePath:    filesWithExt,
// 		CallbackEnv: nil,
// 	}
//
// 	return configs.Load(cfgOption)
// }
