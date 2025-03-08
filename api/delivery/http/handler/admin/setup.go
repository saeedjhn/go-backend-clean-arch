package admin

import (
	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-domain-driven-design/internal/bootstrap"
)

func Setup(_ *bootstrap.Application, e *echo.Echo) {
	// Dependencies

	New().SetRoutes(e)
}
