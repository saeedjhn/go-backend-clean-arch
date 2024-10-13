package taskrouter

import (
	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/api/v1/delivery/http/handler/taskhandler"
	"github.com/saeedjhn/go-backend-clean-arch/api/v1/delivery/http/middleware"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/token"
	"github.com/saeedjhn/go-backend-clean-arch/internal/repository/taskrepository/mysqltask"
	"github.com/saeedjhn/go-backend-clean-arch/internal/service/authservice"
	"github.com/saeedjhn/go-backend-clean-arch/internal/service/taskservice"
	"github.com/saeedjhn/go-backend-clean-arch/internal/validator/taskvalidator"
)

func New(app *bootstrap.Application, group *echo.Group) {
	// Repository
	taskMysql := mysqltask.New(app.MysqlDB)

	// Usecase
	taskSvc := taskservice.New(taskMysql)
	authSvc := authservice.New(app.Config.Auth, token.New())

	// Validator
	validator := taskvalidator.New()

	// Handler
	handler := taskhandler.New(app, validator, taskSvc)

	tasksGroup := group.Group("/tasks")

	protectedRouter := tasksGroup.Group("")
	protectedRouter.Use(middleware.Auth(app.Config.Auth, authSvc))
	{
		protectedRouter.POST("/", handler.Create)
		protectedRouter.GET("/:id", handler.FindOne)
		protectedRouter.GET("/", handler.FindAll)
		protectedRouter.PATCH("/", handler.Update)
		protectedRouter.DELETE("/", handler.Delete)
	}
}
