package healthcheck

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) SetRoutes(e *echo.Echo) {
	group := e.Group("/health-check")
	{
		group.GET("", func(c echo.Context) error {
			return c.JSON(http.StatusOK, echo.Map{
				"message": "everything is good!",
			})
		})
	}
}
