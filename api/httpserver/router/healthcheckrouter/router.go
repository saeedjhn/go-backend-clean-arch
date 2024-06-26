package healthcheckrouter

import (
	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"net/http"
)

func New(_ *bootstrap.Application, e *echo.Group) {
	g := e.Group("/health-check")
	{
		g.GET("/", func(c echo.Context) error {
			return c.JSON(http.StatusOK, echo.Map{
				"message": "everything is good!",
			})
		})
	}
}
