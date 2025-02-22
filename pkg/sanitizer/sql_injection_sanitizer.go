package sanitizer

import (
	"regexp"
)

var sqlInjectionPatterns = []*regexp.Regexp{ //nolint:gochecknoglobals // nothing
	regexp.MustCompile(`--\s*$`),
	regexp.MustCompile(`;`),
	regexp.MustCompile(`/\*.*\*/`),
	regexp.MustCompile(`(?i)\b(exec|execute|sleep|benchmark|xp_)\b`),
	regexp.MustCompile(`(?i)\b(or|and)\s+\d+\s*=\s*\d+\b`),
}

type SQLInjectionSanitizer struct{}

func NewSQLInjectionSanitizer() *SQLInjectionSanitizer {
	return &SQLInjectionSanitizer{}
}

func (s SQLInjectionSanitizer) Sanitize(input string) (string, error) {
	for _, pattern := range sqlInjectionPatterns {
		if pattern.MatchString(input) {
			return "", errSQLInjectDangerous
		}
	}

	return input, nil
}
