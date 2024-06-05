package taskrouter

import (
	"github.com/labstack/echo/v4"
	"go-backend-clean-arch/internal/bootstrap"
	"net/http"
)

func New(app *bootstrap.Application, e *echo.Group) {
	g := e.Group("/tasks")
	{
		g.GET("/", func(c echo.Context) error {
			return c.JSON(http.StatusOK, "-- tasks --")
		})
	}
}
