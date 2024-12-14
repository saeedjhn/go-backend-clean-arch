package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const _HSTSMaxAgeOneYear = 31536000

func Secure() echo.MiddlewareFunc {
	return middleware.SecureWithConfig(middleware.SecureConfig{
		Skipper:               nil,
		CSPReportOnly:         true,                              // dev: true, prod: false
		HSTSPreloadEnabled:    false,                             // HTTP Strict Transport Security, HTTP: false, HTTPS: true
		ReferrerPolicy:        "strict-origin-when-cross-origin", // default: no-referrer
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "true",
		XFrameOptions:         "DENY",
		HSTSMaxAge:            _HSTSMaxAgeOneYear,
		HSTSExcludeSubdomains: true, // Excludes subdomains from the HSTS policy
		ContentSecurityPolicy: "default-src 'self'; script-src 'self'; style-src 'self'; img-src 'self';",
	})
}
