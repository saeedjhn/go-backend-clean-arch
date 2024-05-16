package taskrouter

import (
	"github.com/labstack/echo/v4"
	"go-backend-clean-arch-according-to-go-standards-project-layout/configs"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/db/mysql"
	"net/http"
)

func New(cfg *configs.Config, mysqlDB *mysql.MySqlDB, e *echo.Echo) {
	g := e.Group("tasks")
	{
		g.GET("/", func(c echo.Context) error {
			return c.JSON(http.StatusOK, "-- tasks --")
		})
	}
}
