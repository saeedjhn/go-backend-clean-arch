package suites_test

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/saeedjhn/go-backend-clean-arch/configs"
	setuptest "github.com/saeedjhn/go-backend-clean-arch/test/setup_test"
	"github.com/saeedjhn/go-backend-clean-arch/test/steps/user/doubles"
)

var (
	featuresPath string                  //nolint:gochecknoglobals // nothing
	_myConfig    *configs.Config         //nolint:gochecknoglobals // nothing
	_myTrc       *doubles.DummyTracer    //nolint:gochecknoglobals // nothing
	_configPath  = "testdata/config.yml" //nolint:gochecknoglobals // nothing
)

func TestMain(m *testing.M) {
	log.Println("Setting up resources...")

	featuresPath, _ = filepath.Abs("../features")

	wd, err := os.Getwd()
	if err != nil {
		log.Panicf("error getting current working directory: %v", err)
	}

	if _myConfig, err = setuptest.LoadConfig[*configs.Config](filepath.Join(wd, _configPath)); err != nil {
		log.Panicf("failed to load config: %v", err)
	}

	_myTrc = setupTracer()

	m.Run()
}

func setupTracer() *doubles.DummyTracer {
	return doubles.NewDummyTracer()
}
