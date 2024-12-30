package admin

import (
	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
)

func Setup(app *bootstrap.Application, e *echo.Echo) {
	// Dependencies

	New().SetRoutes(e)
}
