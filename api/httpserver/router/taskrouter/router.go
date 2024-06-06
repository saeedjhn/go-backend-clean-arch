package taskrouter

import (
	"github.com/labstack/echo/v4"
	"go-backend-clean-arch/internal/bootstrap"
	"net/http"
)

func New(_ *bootstrap.Application, group *echo.Group) {
	tasksGroup := group.Group("/tasks")
	{
		tasksGroup.GET("/", func(c echo.Context) error {
			return c.JSON(http.StatusOK, "-- tasks --")
		})
	}
}
