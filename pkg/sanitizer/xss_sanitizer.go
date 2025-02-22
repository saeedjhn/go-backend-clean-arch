package sanitizer

import (
	"github.com/microcosm-cc/bluemonday"
	"regexp"
)

// SanitizeStrictPolicy - bluemonday.XSSStrictPolicy() which can be thought of as equivalent to stripping all HTML
// elements and their attributes as it has nothing on its allowlist.
// An example usage scenario would be blog post titles where HTML tags are not expected at all and if they are
// then the elements and the content of the elements
// should be stripped. This is a very strict policy.
// func sanitizeStrictPolicy(ctx fiber.Ctx) error {
//	return ctx.Next()
// }

// SanitizeUGCPolicy -  bluemonday.XSSUGCPolicy() which allows a broad selection of HTML elements
// and attributes that are safe for userentity generated content.
// Note that this policy does not allow iframes, object, embed, styles, script, etc. An example usage scenario
// would be blog post bodies where a variety of formatting is expected along with the potential for TABLEs and IMGs.
// func sanitizeUGCPolicy(ctx fiber.Ctx) error {
//	return ctx.Next()
// }

type XSSPolicy string

const (
	XSSStrictPolicy    XSSPolicy = "strict_policy"
	XSSUGCPolicy       XSSPolicy = "ugc_policy"
	XSSStripTagsPolicy XSSPolicy = "strip_tags_policy"
)

type XSSSanitizer struct {
	policy *bluemonday.Policy
}

func NewXSSSanitizer(policy XSSPolicy) *XSSSanitizer {
	sanitizer := &XSSSanitizer{}

	return sanitizer.SetPolicy(policy)
}

func (x *XSSSanitizer) SetPolicy(policy XSSPolicy) *XSSSanitizer {
	switch policy {
	case XSSStrictPolicy:
		x.policy = bluemonday.StrictPolicy()
	case XSSUGCPolicy:
		x.policy = bluemonday.UGCPolicy()
	case XSSStripTagsPolicy:
		x.policy = bluemonday.StripTagsPolicy()
	// case NewPolicy:

	//	s.policy = bluemonday.NewPolicy() // TODO - Implement new policy for sanitize
	default:
		x.policy = bluemonday.StrictPolicy()
	}

	return x
}

func (x *XSSSanitizer) Sanitize(input string) (string, error) {
	// regex := regexp.MustCompile(`(?i)javascript\s*:[^;]*(;|$)|[\x00-\x1F\x7F]`)
	regex := regexp.MustCompile(`(?i)javascript\s*:[^;]*(;|$)`)

	if regex.MatchString(input) {
		return "", errXSSAtackDangerous
	}

	sanitized := x.policy.Sanitize(input)

	if sanitized != input {
		return "", errXSSAtackDangerous
	}

	return sanitized, nil
}
