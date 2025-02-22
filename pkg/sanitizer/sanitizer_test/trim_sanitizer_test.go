package sanitizer_test

import (
	"testing"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/sanitizer"
	"github.com/stretchr/testify/assert"
)

//go:generate go test -v -race -count=1 ./...

//go:generate go test -v -race -count=1 -run TestTrimSanitizer_Sanitize
func TestTrimSanitizer_Sanitize(t *testing.T) {
	t.Parallel()

	trimSanitizer := sanitizer.NewTrimSanitizer()

	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"TrimSanitizer_InputWithLeadingSpace_Trimmed", "  hello", "hello"},
		{"TrimSanitizer_InputWithTrailingSpace_Trimmed", "world  ", "world"},
		{"TrimSanitizer_InputWithLeadingAndTrailingSpace_Trimmed", "  test  ", "test"},
		{"TrimSanitizer_InputWithTabs_Trimmed", "\ttext\t", "text"},
		{"TrimSanitizer_InputWithNewLines_Trimmed", "\nhello\n", "hello"},
		{"TrimSanitizer_InputWithMixedSpaces_Trimmed", " \t\n go \n\t ", "go"},
		{"TrimSanitizer_InputWithoutSpaces_Unchanged", "golang", "golang"},
		{"TrimSanitizer_InputWithOnlySpaces_Empty", "     ", ""},
		{"TrimSanitizer_InputWithOnlyTabs_Empty", "\t\t\t", ""},
		{"TrimSanitizer_InputWithOnlyNewLines_Empty", "\n\n\n", ""},
		{"TrimSanitizer_InputWithSpacesBetweenWords_Unchanged", "foo bar", "foo bar"},
		{"TrimSanitizer_InputWithMultipleSpacesBetweenWords_Unchanged", "foo    bar", "foo    bar"},
		{"TrimSanitizer_InputWithUnicodeSpaces_Trimmed", " 你好 ", "你好"},
		{"TrimSanitizer_InputWithSpecialCharacters_Trimmed", " @test@ ", "@test@"},
		{"TrimSanitizer_EmptyInput_Empty", "", ""},
		{"TrimSanitizer_InputWithNonBreakingSpace_Trimmed", "\u00A0hello\u00A0", "hello"},
		{"TrimSanitizer_InputWithCarriageReturn_Trimmed", "\rtest\r", "test"},
		{"TrimSanitizer_InputWithFormFeed_Trimmed", "\ftest\f", "test"},
		{"TrimSanitizer_InputWithVerticalTab_Trimmed", "\vtest\v", "test"},
		{"TrimSanitizer_InputWithMixedWhitespaceCharacters_Trimmed", "\t \n\rfoo\v\f", "foo"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result, _ := trimSanitizer.Sanitize(tc.input)

			assert.Equal(t, tc.expected, result)
		})
	}
}
