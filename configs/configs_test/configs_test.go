package configs_test

import (
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/saeedjhn/go-backend-clean-arch/configs"
)

const testdataPath = "testdata"
const fileExtension = "yml"

var (
	_configFlag  string //nolint:gochecknoglobals // nothing
	_fileExtFlag string //nolint:gochecknoglobals // nothing
)

//nolint:gochecknoinits // nothing
func init() {
	flag.StringVar(&_configFlag, "config", testdataPath, "config path, e.g., -conf configs")
	flag.StringVar(&_fileExtFlag, "extension", fileExtension, "The files extension")
}

func Test_CollectFilesWithExt_ValidDirectoryAndExtension_ReturnsFileList(t *testing.T) {
	workingDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting current working directory: %v", err)
	}

	configPath := filepath.Join(workingDir, testdataPath)

	t.Log(configs.CollectFilesWithExt(configPath, "yml"))
}

//go:generate go test -v -run Test_Load_ValidConfigFilesAndOptions_ReturnsConfigSuccessfully
//go:generate go test -v -run Test_Load_ValidConfigFilesAndOptions_ReturnsConfigSuccessfully -config=testdata -extension=yml
func Test_Load_ValidConfigFilesAndOptions_ReturnsConfigSuccessfully(t *testing.T) {
	workingDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting current working directory: %v", err)
	}

	configPath := filepath.Join(workingDir, _configFlag)

	filesWithExt, err := configs.CollectFilesWithExt(
		configPath,
		_fileExtFlag,
	)
	if err != nil {
		t.Fatalf(
			"Unexpected error while loading configuration files from directory: %s. Error: %v",
			filepath.Join(workingDir, _configFlag),
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
		t.Fatalf("Error loading configuration with option '%v': %v", cfgOption, err)
	}

	t.Logf("%#v", config)
}
