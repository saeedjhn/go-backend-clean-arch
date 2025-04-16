package healthz

import (
	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
)

func Setup(_ *bootstrap.Application, e *echo.Echo) {
	// Dependencies

	handler := New()
	handler.SetRoutes(e)
}
