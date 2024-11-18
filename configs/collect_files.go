package configs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CollectFilesWithExt(
	baseDir,
	subDir,
	ext string,
) ([]string, error) {
	dirPath := filepath.Join(baseDir, subDir)

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %w", err)
	}

	var filesList []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), fmt.Sprintf(".%s", ext)) {
			filesList = append(filesList, fmt.Sprintf("%s/%s", dirPath, entry.Name()))
		}
	}

	return filesList, nil
}
