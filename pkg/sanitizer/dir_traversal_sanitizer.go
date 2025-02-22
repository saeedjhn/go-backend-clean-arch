package sanitizer

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

type DirTraversalSanitizer struct {
	basePath string
}

func NewDirTraversalSanitizer(basePath string) *DirTraversalSanitizer {
	return &DirTraversalSanitizer{basePath: basePath}
}

func (t *DirTraversalSanitizer) Sanitize(input string) (string, error) {
	if filepath.IsAbs(input) {
		return "", errDirectoryTraversalDangerous
	}

	cleanPath := input
	if len(input) != 0 {
		cleanPath = filepath.Clean(input)
	}

	fullPath := filepath.Join(t.basePath, cleanPath)

	realPath, err := filepath.EvalSymlinks(fullPath)
	if err != nil {
		return "", fmt.Errorf("invalid path: %w", err)
	}

	if !strings.HasPrefix(realPath, t.basePath) {
		return "", fmt.Errorf("access denied: %s", realPath)
	}
	log.Println(input, cleanPath)

	return cleanPath, nil
}
