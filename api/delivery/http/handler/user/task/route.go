package task

import (
	"github.com/labstack/echo/v4"
	mymiddleware "github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/middleware"
)

func (h *Handler) SetRoutes(e *echo.Echo) {
	group := e.Group("/users")

	group.Use(mymiddleware.Auth(h.authIntr))
	{
		group.POST("/:id/tasks", h.Create, mymiddleware.CheckIsValidUserID)
		group.GET("/:id/tasks", h.Tasks, mymiddleware.CheckIsValidUserID)
	}
}
