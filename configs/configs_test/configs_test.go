package configs_test

import (
	"os"
	"testing"

	"github.com/saeedjhn/go-backend-clean-arch/configs"
)

func TestSource(t *testing.T) {
	wDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting current working directory: %v", err)
	}

	t.Log(configs.CollectFilesWithExt(wDir, "", "yml"))
}
