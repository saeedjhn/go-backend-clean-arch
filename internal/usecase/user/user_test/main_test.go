package user_test

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/saeedjhn/go-backend-clean-arch/configs"
)

const (
	_testdataPath  = "testdata"
	_fileExtension = "yml"
)

var (
	errInvalidInput              = errors.New("invalid input")
	errUserNotFound              = errors.New("user not found")
	errAccessTokenCreationFailed = errors.New("failed create token")
	errUnexpected                = errors.New("unexpected error")
)

func TestMain(m *testing.M) {
	log.Println("Setting up resources...")

	// var _configPath string
	// flag.StringVar(&_configPath, "config", "config.json", "Path to the config file")
	// If your tests depend on command line flags, you will need to call flag.Parse() manually.
	flag.Parse()

	// DB connection

	code := m.Run()

	log.Println("Tearing down resources...")
	// DB close

	os.Exit(code)
}

func getConfig() (*configs.Config, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return &configs.Config{}, fmt.Errorf("error getting current working directory: %w", err)
	}

	configPath := filepath.Join(workingDir, _testdataPath)

	filesWithExt, err := configs.CollectFilesWithExt(
		configPath,
		_fileExtension,
	)
	if err != nil {
		return &configs.Config{}, fmt.Errorf(
			"unexpected error while loading configuration files from directory: %s. Error: %w",
			filepath.Join(workingDir, _testdataPath),
			err,
		)
	}

	cfgOption := configs.Option{
		Prefix:      "",
		Delimiter:   "",
		Separator:   "",
		FilePath:    filesWithExt,
		CallbackEnv: nil,
	}

	config, err := configs.Load(cfgOption)
	if err != nil {
		return &configs.Config{}, fmt.Errorf(
			"error loading configuration with option '%v': %w", cfgOption, err,
		)
	}

	return config, nil
}
