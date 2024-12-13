package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const _HSTSMaxAgeOneYear = 31536000

func Secure() echo.MiddlewareFunc {
	return middleware.SecureWithConfig(middleware.SecureConfig{
		Skipper:               nil,
		CSPReportOnly:         false,
		HSTSPreloadEnabled:    false,
		ReferrerPolicy:        "",
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "true",
		XFrameOptions:         "DENY",
		HSTSMaxAge:            _HSTSMaxAgeOneYear,
		HSTSExcludeSubdomains: true, // Excludes subdomains from the HSTS policy
		ContentSecurityPolicy: "default-src 'self'; script-src 'self'; style-src 'self'; img-src 'self';",
	})
}
