package healthz

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) SetRoutes(e *echo.Echo) {
	group := e.Group("/healthz")
	{
		// This endpoint checks whether the service is still alive or not.
		group.GET("/liveness", func(c echo.Context) error {
			return c.NoContent(http.StatusOK)
			// return c.JSON(http.StatusOK, echo.Map{
			// 	"message": "everything is good!",
			// })
		})

		// This checks whether the service is ready to accept the request.
		group.GET("/readiness", func(c echo.Context) error {
			return c.NoContent(http.StatusOK)
		})
	}
}
