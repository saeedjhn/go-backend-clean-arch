package setuptest_test

import (
	"os"
	"path/filepath"
	"testing"

	setuptest "github.com/saeedjhn/go-backend-clean-arch/test/setup_test"
)

const (
	_configPath = "testdata/config.yml"
)

type AppConfig struct {
	Env  string `mapstructure:"env"`
	Port int    `mapstructure:"port"`
}

func Test_LoadConfig(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("error getting current working directory: %v", err)
	}

	config, err := setuptest.LoadConfig[AppConfig](filepath.Join(wd, _configPath))
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	t.Logf("Loaded Config: %+v\n", config)
}
