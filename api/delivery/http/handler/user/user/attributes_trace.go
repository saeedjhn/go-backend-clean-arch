package user

import (
	"maps"

	"github.com/labstack/echo/v4"
)

func attributes(
	ctx echo.Context,
	extraAttrs ...map[string]interface{}, //nolint:unparam // nothing
) map[string]interface{} {
	defaultAttrs := map[string]interface{}{
		"http.method":     ctx.Request().Method,
		"http.host":       ctx.Request().Host,
		"http.route":      ctx.Path(),
		"http.client_ip":  ctx.RealIP(),
		"http.scheme":     ctx.Scheme(),
		"http.user_agent": ctx.Request().UserAgent(),
	}

	if len(extraAttrs) != 0 {
		for _, attr := range extraAttrs {
			maps.Copy(defaultAttrs, attr)
		}
	}

	return defaultAttrs
}
