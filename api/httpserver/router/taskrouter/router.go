package taskrouter

import (
	"github.com/labstack/echo/v4"
	"go-backend-clean-arch/api/httpserver/handler/taskhandler"
	"go-backend-clean-arch/api/httpserver/middleware"
	"go-backend-clean-arch/internal/bootstrap"
	"go-backend-clean-arch/internal/infrastructure/token"
	"go-backend-clean-arch/internal/repository/taskrepository/mysqltask"
	"go-backend-clean-arch/internal/usecase/authusecase"
	"go-backend-clean-arch/internal/usecase/taskusecase"
	"go-backend-clean-arch/internal/validator/taskvalidator"
)

func New(app *bootstrap.Application, group *echo.Group) {
	// Repository
	taskMysql := mysqltask.New(app.MysqlDB)

	// Usecase
	taskCase := taskusecase.New(taskMysql)
	authCase := authusecase.New(app.Config.Auth, token.New())

	// Validator
	validator := taskvalidator.New()

	// Handler
	handler := taskhandler.New(app, validator, taskCase)

	tasksGroup := group.Group("/tasks")

	protectedRouter := tasksGroup.Group("")
	protectedRouter.Use(middleware.Auth(app.Config.Auth, authCase))
	{
		protectedRouter.POST("/", handler.Create)
		protectedRouter.GET("/:id", handler.FindOne)
		protectedRouter.GET("/", handler.FindAll)
		protectedRouter.PATCH("/", handler.Update)
		protectedRouter.DELETE("/", handler.Delete)
	}
}
