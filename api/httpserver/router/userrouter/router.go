package userrouter

import (
	"github.com/labstack/echo/v4"
	"go-backend-clean-arch/api/httpserver/handler/userhandler"
	"go-backend-clean-arch/internal/bootstrap"
	"go-backend-clean-arch/internal/gateway/taskinggateway"
	"go-backend-clean-arch/internal/infrastructure/token"
	"go-backend-clean-arch/internal/repository/taskrepository/mysqltask"
	"go-backend-clean-arch/internal/repository/userrespository/mysqluser"
	"go-backend-clean-arch/internal/usecase/taskusecase"
	"go-backend-clean-arch/internal/usecase/userusecase"
	"go-backend-clean-arch/internal/validator/uservalidator"
)

func New(
	app *bootstrap.Application,
	group *echo.Group,
) {
	// dependencies
	t := token.New()

	// Repository
	taskMysql := mysqltask.New(app.MysqlDB)
	userMysql := mysqluser.New(app.MysqlDB)

	// Repository & Usecase
	taskCase := taskusecase.New(taskMysql)

	// Service-oriented - no depends on useCases - ( userusecase -> taskgateway -> taskusecase )
	taskGate := taskinggateway.New(taskCase)

	// Repository & Usecase
	userCase := userusecase.New(
		app.Config, taskGate, userMysql, t,
	)

	// Validator
	validator := uservalidator.New()

	// Handler
	handler := userhandler.New(app, validator, userCase)

	usersGroup := group.Group("/users")

	publicRouter := usersGroup.Group("/auth")
	{
		publicRouter.POST("/register", handler.Register)
		publicRouter.POST("/login", handler.Login)
	}

	protectedRouter := usersGroup.Group("")

	//protectedRouter.Use(middleware.Auth())
	{
		protectedRouter.GET("/task-list", handler.TaskList)
	}
}
