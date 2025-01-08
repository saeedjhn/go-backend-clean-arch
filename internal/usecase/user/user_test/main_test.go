package user_test

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/saeedjhn/go-backend-clean-arch/configs"
	setuptest "github.com/saeedjhn/go-backend-clean-arch/test/setup_test"
)

var (
	_myDBConfig *configs.Config
	_configPath = "testdata/config.yml"

	errInvalidInput              = errors.New("invalid input")
	errUserNotFound              = errors.New("user not found")
	errAccessTokenCreationFailed = errors.New("failed create token")
	errUnexpected                = errors.New("unexpected error")
)

func TestMain(m *testing.M) {
	log.Println("Setting up resources...")

	wd, err := os.Getwd()
	if err != nil {
		log.Panicf("error getting current working directory: %v", err)
	}

	_myDBConfig, err = setuptest.LoadConfig[*configs.Config](filepath.Join(wd, _configPath))
	if err != nil {
		log.Panicf("failed to load config: %v", err)
	}

	m.Run()
}
