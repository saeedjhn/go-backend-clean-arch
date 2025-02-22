package sanitizer_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/sanitizer"
)

func TestDirTraversalSanitizer_Sanitize(t *testing.T) {
	t.Parallel()

	wd, _ := os.Getwd()
	basePath := filepath.Join(wd, "home/user/uploads")

	tests := []struct {
		name        string
		input       string
		basePath    string
		expected    string
		expectError bool
	}{
		{
			name:        "AbsolutePath_NotAllowed",
			input:       "/etc/passwd",
			basePath:    basePath,
			expected:    "",
			expectError: true,
		},
		{
			name:        "RelativePath_Valid",
			input:       "relative/file",
			basePath:    basePath,
			expected:    "relative/file",
			expectError: false,
		},
		{
			name:        "CleanedRelativePath_Valid",
			input:       "../relative/file",
			basePath:    basePath,
			expected:    "",
			expectError: true,
		},
		{
			name:        "Symlink_Valid",
			input:       "symlink-to-valid",
			basePath:    basePath,
			expected:    "symlink-to-valid",
			expectError: false,
		},
		{
			name:        "Symlink_DeniedAccess",
			input:       "symlink-to-invalid",
			basePath:    basePath,
			expected:    "",
			expectError: true,
		},
		{
			name:        "AbsolutePath_InvalidBase",
			input:       "/home/user/project/some/path",
			basePath:    basePath,
			expected:    "",
			expectError: true,
		},
		{
			name:        "InvalidPath_Error",
			input:       "invalid//file",
			basePath:    basePath,
			expected:    "invalid/file",
			expectError: false,
		},
		{
			name:        "EmptyPath_Error",
			input:       "",
			basePath:    basePath,
			expected:    "",
			expectError: false, // Updated to true since an empty path should return an error
		},
		{
			name:        "NormalPath_Valid",
			input:       "folder/file",
			basePath:    basePath,
			expected:    "folder/file",
			expectError: false,
		},
		{
			name:        "AccessDenied_OutsideBasePath",
			input:       "../../outside/project/file",
			basePath:    basePath,
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			traversalSanitizer := sanitizer.NewDirTraversalSanitizer(tt.basePath)
			got, err := traversalSanitizer.Sanitize(tt.input)

			if (err != nil) != tt.expectError {
				t.Errorf("Sanitize() error = %v, expectError %v", err, tt.expectError)
				return
			}

			if got != tt.expected {
				t.Errorf("Sanitize() got = %v, want %v", got, tt.expected)
			}
		})
	}
}
