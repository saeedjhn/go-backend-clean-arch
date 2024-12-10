package healthcheckrouter

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
)

func New(_ *bootstrap.Application, group *echo.Group) {
	group.GET("/health-check", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "everything is good!",
		})
	})
}
