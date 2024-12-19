package task

import (
	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/handler/task"
	"github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/middleware"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	vld "github.com/saeedjhn/go-backend-clean-arch/internal/validator/task"
)

func New(app *bootstrap.Application, group *echo.Group) {
	// Validator
	validator := vld.New()

	// Handler
	handler := task.New(
		app,
		app.Trc,
		validator,
		app.Usecase.TaskIntr,
	)

	tasksGroup := group.Group("/tasks")

	protectedRouter := tasksGroup.Group("")

	protectedRouter.Use(middleware.Auth(app.Usecase.AuthIntr))
	{
		protectedRouter.POST("/", handler.Create)
		protectedRouter.GET("/:id", handler.FindOne)
		protectedRouter.GET("/", handler.FindAll)
		protectedRouter.PATCH("/", handler.Update)
		protectedRouter.DELETE("/", handler.Delete)
	}
}
