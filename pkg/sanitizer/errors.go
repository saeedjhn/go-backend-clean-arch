package sanitizer

import (
	"errors"
)

var (
	errNonPointer = errors.New("expected a pointer to a struct, but received a non-pointer value")

	errDirectoryTraversalDangerous = errors.New(
		"[Directory Traversal attack]: potential directory traversal attack detected")
	errSQLInjectDangerous = errors.New("[SQL injection]: input contains a dangerous SQL Injection pattern")
	errXSSAtackDangerous  = errors.New("[XSS attack]: potential XSS attack detected in input")
)
