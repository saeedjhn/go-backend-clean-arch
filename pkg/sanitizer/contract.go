package sanitizer

type SanitizationStrategy interface {
	Sanitize(input string) (string, error)
}
