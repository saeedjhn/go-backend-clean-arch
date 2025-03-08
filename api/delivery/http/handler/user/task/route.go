package task

import (
	"github.com/labstack/echo/v4"
	mymiddleware "github.com/saeedjhn/go-domain-driven-design/api/delivery/http/middleware"
)

func (h *Handler) SetRoutes(e *echo.Echo) {
	group := e.Group("/users")

	group.Use(mymiddleware.Authentication(h.authIntr))
	{
		group.POST("/:id/tasks", h.Create, mymiddleware.CheckIsValidUserID)
		group.GET("/:id/tasks", h.Tasks, mymiddleware.CheckIsValidUserID)
	}
}
