package sanitizer_test

import (
	"testing"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/sanitizer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:generate go test -v -race -count=1 ./...

//go:generate go test -v -race -count=1 -run TestXSSSanitizer_Sanitize
func TestXSSSanitizer_Sanitize(t *testing.T) {
	tests := []struct {
		name      string
		policy    sanitizer.XSSPolicy
		input     string
		expectErr bool
	}{
		{"StrictPolicy_ValidHTML_removed", sanitizer.XSSStrictPolicy, "<b>Bold</b>", true},
		{"StrictPolicy_ValidHTMLWithUGCPolicy_Allowed", sanitizer.XSSStrictPolicy, "<b>Bold</b>", true},
		{"StrictPolicy_ScriptTag_Removed", sanitizer.XSSStrictPolicy, "<script>alert(1)</script>", true},
		{"StrictPolicy_JavascriptURL_Removed", sanitizer.XSSStrictPolicy, "javascript:alert(1)", true},
		{
			"StrictPolicy_OnClickAttribute_Removed",
			sanitizer.XSSStrictPolicy,
			"<a href='#' onclick='alert(1)'>Click</a>",
			true,
		},
		{"StrictPolicy_SVGInjection_Removed", sanitizer.XSSStrictPolicy, "<svg><script>alert(1)</script></svg>", true},

		{"UGCPolicy_ValidHTML_Allowed", sanitizer.XSSUGCPolicy, "<b>Bold</b>", false},
		{"UGCPolicy_ScriptTag_Removed", sanitizer.XSSUGCPolicy, "<script>alert(1)</script>", true},
		{"UGCPolicy_JavascriptURL_Removed", sanitizer.XSSUGCPolicy, "javascript:alert(1)", true},
		{"UGCPolicy_OnClickAttribute_Removed", sanitizer.XSSUGCPolicy, "<a href='#' onclick='alert(1)'>Click</a>", true},
		{"UGCPolicy_AllowedTags_Retained", sanitizer.XSSUGCPolicy, "<p>Hello</p>", false},

		{"StrictPolicy_EmptyString_NoError", sanitizer.XSSStrictPolicy, "", false},
		{"StrictPolicy_Whitespace_NoError", sanitizer.XSSStrictPolicy, "  ", false},
		{"StrictPolicy_SafeText_NoError", sanitizer.XSSStrictPolicy, "Hello World!", false},
		{
			"StrictPolicy_EncodedScript_Removed",
			sanitizer.XSSStrictPolicy,
			"&#x3C;script&#x3E;alert(1)&#x3C;/script&#x3E;",
			true,
		},
		{
			"StrictPolicy_HexEncodedJS_Removed",
			sanitizer.XSSStrictPolicy,
			"&#x6a;&#x61;&#x76;&#x61;&#x73;&#x63;&#x72;&#x69;&#x70;&#x74;:alert(1)",
			true,
		},

		{"UGCPolicy_EmptyString_NoError", sanitizer.XSSUGCPolicy, "", false},
		{"UGCPolicy_Whitespace_NoError", sanitizer.XSSUGCPolicy, "  ", false},
		{"UGCPolicy_SafeText_NoError", sanitizer.XSSUGCPolicy, "Hello World!", false},
		{"UGCPolicy_AllowedHTML_NoError", sanitizer.XSSUGCPolicy, "<i>Italic</i>", false},
		{"UGCPolicy_EncodedScript_Removed", sanitizer.XSSUGCPolicy, "&#x3C;script&#x3E;alert(1)&#x3C;/script&#x3E;", true},

		{"StrictPolicy_ScriptTag_Removed", sanitizer.XSSStrictPolicy, "<script>alert(1)</script>", true},
		{"StrictPolicy_JavascriptURL_Removed", sanitizer.XSSStrictPolicy, "javascript:alert(1)", true},
		{
			"StrictPolicy_OnClickAttribute_Removed",
			sanitizer.XSSStrictPolicy,
			"<a href='#' onclick='alert(1)'>Click</a>",
			true,
		},
		{"StrictPolicy_SVGInjection_Removed", sanitizer.XSSStrictPolicy, "<svg><script>alert(1)</script></svg>", true},
		{
			"StrictPolicy_IFRAMEInjection_Removed",
			sanitizer.XSSStrictPolicy,
			"<iframe src='javascript:alert(1)'></iframe>",
			true,
		},
		{"StrictPolicy_ObjectTag_Removed", sanitizer.XSSStrictPolicy, "<object data='javascript:alert(1)'></object>", true},
		{
			"StrictPolicy_MetaRefresh_Removed",
			sanitizer.XSSStrictPolicy,
			"<meta http-equiv='refresh' content='0;url=javascript:alert(1)'>",
			true,
		},
		{
			"StrictPolicy_StyleExpression_Removed",
			sanitizer.XSSStrictPolicy,
			"<div style='expression(alert(1))'>Test</div>",
			true,
		},
		{
			"StrictPolicy_BackgroundImageJS_Removed",
			sanitizer.XSSStrictPolicy,
			"<div style='background:url(javascript:alert(1))'>Test</div>",
			true,
		},

		{"UGCPolicy_ValidHTML_Allowed", sanitizer.XSSUGCPolicy, "<b>Bold</b>", false},
		{"UGCPolicy_ScriptTag_Removed", sanitizer.XSSUGCPolicy, "<script>alert(1)</script>", true},
		{"UGCPolicy_JavascriptURL_Removed", sanitizer.XSSUGCPolicy, "javascript:alert(1)", true},
		{"UGCPolicy_OnClickAttribute_Removed", sanitizer.XSSUGCPolicy, "<a href='#' onclick='alert(1)'>Click</a>", true},
		{"UGCPolicy_AllowedTags_Retained", sanitizer.XSSUGCPolicy, "<p>Hello</p>", false},
		{"UGCPolicy_IFRAMEInjection_Removed", sanitizer.XSSUGCPolicy, "<iframe src='javascript:alert(1)'></iframe>", true},
		{"UGCPolicy_ObjectTag_Removed", sanitizer.XSSUGCPolicy, "<object data='javascript:alert(1)'></object>", true},
		{
			"UGCPolicy_MetaRefresh_Removed",
			sanitizer.XSSUGCPolicy,
			"<meta http-equiv='refresh' content='0;url=javascript:alert(1)'>",
			true,
		},
		{"UGCPolicy_StyleExpression_Removed", sanitizer.XSSUGCPolicy, "<div style='expression(alert(1))'>Test</div>", true},
		{
			"UGCPolicy_BackgroundImageJS_Removed",
			sanitizer.XSSUGCPolicy,
			"<div style='background:url(javascript:alert(1))'>Test</div>",
			true,
		},

		{"StrictPolicy_EmptyString_NoError", sanitizer.XSSStrictPolicy, "", false},
		{"StrictPolicy_Whitespace_NoError", sanitizer.XSSStrictPolicy, "  ", false},
		{"StrictPolicy_SafeText_NoError", sanitizer.XSSStrictPolicy, "Hello World!", false},
		{
			"StrictPolicy_EncodedScript_Removed",
			sanitizer.XSSStrictPolicy,
			"&#x3C;script&#x3E;alert(1)&#x3C;/script&#x3E;",
			true,
		},
		{
			"StrictPolicy_HexEncodedJS_Removed",
			sanitizer.XSSStrictPolicy,
			"&#x6a;&#x61;&#x76;&#x61;&#x73;&#x63;&#x72;&#x69;&#x70;&#x74;:alert(1)",
			true,
		},
		{
			"StrictPolicy_DataURI_Removed",
			sanitizer.XSSStrictPolicy,
			"<img src='data:text/html;base64,PHNjcmlwdD5hbGVydCgxKTwvc2NyaXB0Pg=='>",
			true,
		},
		{"StrictPolicy_MalformedScript_Removed", sanitizer.XSSStrictPolicy, "<scr<script>ipt>alert(1)</script>", true},
		{"StrictPolicy_UnclosedTag_Removed", sanitizer.XSSStrictPolicy, "<script>alert(1)", true},

		{"UGCPolicy_EmptyString_NoError", sanitizer.XSSUGCPolicy, "", false},
		{"UGCPolicy_Whitespace_NoError", sanitizer.XSSUGCPolicy, "  ", false},
		{"UGCPolicy_SafeText_NoError", sanitizer.XSSUGCPolicy, "Hello World!", false},
		{"UGCPolicy_AllowedHTML_NoError", sanitizer.XSSUGCPolicy, "<i>Italic</i>", false},
		{"UGCPolicy_EncodedScript_Removed", sanitizer.XSSUGCPolicy, "&#x3C;script&#x3E;alert(1)&#x3C;/script&#x3E;", true},
		{
			"UGCPolicy_DataURI_Removed",
			sanitizer.XSSUGCPolicy,
			"<img src='data:text/html;base64,PHNjcmlwdD5hbGVydCgxKTwvc2NyaXB0Pg=='>",
			true,
		},
		{"UGCPolicy_MalformedScript_Removed", sanitizer.XSSUGCPolicy, "<scr<script>ipt>alert(1)</script>", true},
		{"UGCPolicy_UnclosedTag_Removed", sanitizer.XSSUGCPolicy, "<script>alert(1)", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			xssSanitizer := sanitizer.NewXSSSanitizer(tc.policy)
			result, err := xssSanitizer.Sanitize(tc.input)

			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.input, result)
			}
		})
	}
}

// func TestXSSSanitizer_Sanitize(t *testing.T) {
// 	t.Parallel()
//
// 	tests := []struct {
// 		name     string
// 		policy   sanitizer.XSSPolicy
// 		input    string
// 		expected string
// 	}{
// 		{
// 			name:     "Sanitize_XSSStringWithsanitizer.XSSStrictPolicy_ShouldSanitizeCorrectly",
// 			policy:   sanitizer.XSSStrictPolicy,
// 			input:    `<script>alert('XSS');</script>`,
// 			expected: ``,
// 		},
// 		{
// 			name:     "Sanitize_XSSStringWithsanitizer.XSSUGCPolicy_ShouldSanitizeCorrectly",
// 			policy:   sanitizer.XSSUGCPolicy,
// 			input:    `<script>alert('XSS');</script>`,
// 			expected: ``,
// 		},
// 		{
// 			name:     "Sanitize_XSSStringWithStripTagsPolicy_ShouldSanitizeCorrectly",
// 			policy:   sanitizer.XSSStripTagsPolicy,
// 			input:    `<script>alert('XSS');</script>`,
// 			expected: ``,
// 		},
// 		{
// 			name:     "Sanitize_XSSStringWithInvalidPolicy_ShouldSanitizeWithsanitizer.XSSStrictPolicy",
// 			policy:   "invalid_policy",
// 			input:    `<script>alert('XSS');</script>`,
// 			expected: ``,
// 		},
// 		{
// 			name:     "Sanitize_EmptyString_ShouldReturnEmptyString",
// 			policy:   sanitizer.XSSStrictPolicy,
// 			input:    ``,
// 			expected: ``,
// 		},
// 		{
// 			name:     "Sanitize_StringWithoutXSS_ShouldReturnSameString",
// 			policy:   sanitizer.XSSStrictPolicy,
// 			input:    `Hello World!`,
// 			expected: `Hello World!`,
// 		},
// 		{
// 			name:     "Sanitize_JavaScriptPseudoProtocol_ShouldRemoveJavascriptPseudoProtocol",
// 			policy:   sanitizer.XSSStrictPolicy,
// 			input:    `javascript:alert('XSS')`,
// 			expected: ``,
// 		},
// 		{
// 			name:     "Sanitize_ValidHTML_ShouldNotAlterContent",
// 			policy:   sanitizer.XSSStrictPolicy,
// 			input:    `<b>Bold Text</b>`,
// 			expected: `Bold Text`,
// 		},
// 		{
// 			name:     "Sanitize_HTMLWithScriptTags_ShouldRemoveScriptTags",
// 			policy:   sanitizer.XSSStrictPolicy,
// 			input:    `<html><script>alert('XSS');</script><body>Content</body></html>`,
// 			expected: `Content`,
// 		},
// 		{
// 			name:     "Sanitize_JavaScriptInHref_ShouldRemoveJavaScriptInHref",
// 			policy:   sanitizer.XSSStrictPolicy,
// 			input:    `<a href="javascript:alert('XSS')">Click me</a>`,
// 			expected: ``,
// 		},
// 		{
// 			name:     "Sanitize_ValidURLInHref_ShouldNotAlterHref",
// 			policy:   sanitizer.XSSStrictPolicy,
// 			input:    `<a href="https://example.com">Visit Example</a>`,
// 			expected: `Visit Example`,
// 		},
// 		{
// 			name:     "Sanitize_ValidURLInHrefWithUGCPolicy_ShouldAlterHref",
// 			policy:   sanitizer.XSSUGCPolicy,
// 			input:    `<a href="https://example.com">Visit Example</a>`,
// 			expected: "<a href=\"https://example.com\" rel=\"nofollow\">Visit Example</a>",
// 		},
// 		{
// 			name:     "Sanitize_EncodedXSS_ShouldSanitizeEncodedXSSCorrectly",
// 			policy:   sanitizer.XSSStrictPolicy,
// 			input:    `&lt;script&gt;alert(&#39;XSS&#39;);&lt;/script&gt;`,
// 			expected: `&lt;script&gt;alert(&#39;XSS&#39;);&lt;/script&gt;`,
// 		},
// 		{
// 			name:     "Sanitize_CSSInjection_ShouldRemoveCSSInjection",
// 			policy:   sanitizer.XSSStrictPolicy,
// 			input:    `<div style="background-image: url(javascript:alert('XSS'));">Content</div>`,
// 			expected: `Content`,
// 		},
// 		{
// 			name:     "Sanitize_MultipleXSSAttacks_ShouldSanitizeAllInstances",
// 			policy:   sanitizer.XSSStrictPolicy,
// 			input:    `<script>alert('XSS');</script><script>alert('XSS');</script>`,
// 			expected: ``,
// 		},
// 		{
// 			name:     "Sanitize_HTMLWithEventAttributes_ShouldRemoveEventAttributes",
// 			policy:   sanitizer.XSSStrictPolicy,
// 			input:    `<div onmouseover="alert('XSS')">Hover me</div>`,
// 			expected: `Hover me`,
// 		},
// 		{
// 			name:     "Sanitize_MultipleValidStrings_ShouldSanitizeEachCorrectly",
// 			policy:   sanitizer.XSSUGCPolicy,
// 			input:    `<div>Hello <script>alert('XSS');</script> World</div>`,
// 			expected: `<div>Hello  World</div>`,
// 		},
// 		{
// 			name:     "Sanitize_WhitespaceAroundJavascript_ShouldRemoveJavascript",
// 			policy:   sanitizer.XSSStrictPolicy,
// 			input:    `javascript :alert('XSS')`,
// 			expected: ``,
// 		},
// 		{
// 			name:     "Sanitize_JavascriptWithSpaces_ShouldRemoveJavascript",
// 			policy:   sanitizer.XSSStrictPolicy,
// 			input:    `javascript   :alert('XSS')`,
// 			expected: ``,
// 		},
// 		{
// 			name:     "Sanitize_JavascriptWithMultipleSpaces_ShouldRemoveJavascript",
// 			policy:   sanitizer.XSSStrictPolicy,
// 			input:    ` javascript : alert ('XSS') `,
// 			expected: ` `,
// 		},
// 		{
// 			name:     "Sanitize_JavascriptWithNewlines_ShouldRemoveJavascript",
// 			policy:   sanitizer.XSSStrictPolicy,
// 			input:    "javascript\n:alert('XSS')",
// 			expected: ``,
// 		},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
//
// 			xssSanitizer := sanitizer.NewXSSSanitizer(tt.policy)
// 			sanitized := xssSanitizer.Sanitize(tt.input)
//
// 			assert.Equal(t, tt.expected, sanitized)
// 		})
// 	}
// }
