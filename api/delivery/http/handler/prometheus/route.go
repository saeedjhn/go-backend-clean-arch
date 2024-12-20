package prometheus

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
)

func (h *Handler) SetRoutes(e *echo.Echo) {
	group := e.Group("/metrics")
	{
		group.GET("", echoprometheus.NewHandler())
	}
}
