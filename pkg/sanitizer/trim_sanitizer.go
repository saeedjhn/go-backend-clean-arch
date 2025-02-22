package sanitizer

import "strings"

type TrimSanitizer struct{}

func NewTrimSanitizer() *TrimSanitizer {
	return &TrimSanitizer{}
}

func (t TrimSanitizer) Sanitize(input string) (string, error) {
	return strings.TrimSpace(input), nil
}
