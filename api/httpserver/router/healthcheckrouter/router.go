package healthcheckrouter

import (
	"github.com/labstack/echo/v4"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/bootstrap"
	"net/http"
)

func New(app *bootstrap.Application, e *echo.Echo) {
	g := e.Group("/health-check")
	{
		g.GET("/", func(c echo.Context) error {
			return c.JSON(http.StatusOK, echo.Map{
				"message": "everything is good!",
			})
		})
	}
}
