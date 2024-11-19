package configs_test

import (
	"os"
	"testing"

	"github.com/saeedjhn/go-backend-clean-arch/configs"
)

func TestCollectFilesWithExt(t *testing.T) {
	workingDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting current working directory: %v", err)
	}

	t.Log(configs.CollectFilesWithExt(workingDir, "yml"))
}
