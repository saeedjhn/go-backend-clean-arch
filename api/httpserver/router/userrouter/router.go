package userrouter

import (
	"github.com/labstack/echo/v4"
	"go-backend-clean-arch/api/httpserver/handler/userhandler"
	"go-backend-clean-arch/api/httpserver/middleware"
	"go-backend-clean-arch/internal/bootstrap"
	"go-backend-clean-arch/internal/infrastructure/token"
	"go-backend-clean-arch/internal/repository/taskrepository/mysqltask"
	"go-backend-clean-arch/internal/repository/userrespository/mysqluser"
	"go-backend-clean-arch/internal/service/authservice"
	"go-backend-clean-arch/internal/service/taskservice"
	"go-backend-clean-arch/internal/service/userservice"
	"go-backend-clean-arch/internal/validator/uservalidator"
)

func New(
	app *bootstrap.Application,
	group *echo.Group,
) {
	// Repository
	taskMysql := mysqltask.New(app.MysqlDB)
	userMysql := mysqluser.New(app.MysqlDB)

	// Usecase
	taskCase := taskservice.New(taskMysql)
	authCase := authservice.New(app.Config.Auth, token.New())

	// Service-oriented - inject service to another service
	// Repository & Usecase
	userCase := userservice.New(
		app.Config, authCase, taskCase, userMysql,
	)

	// Validator
	validator := uservalidator.New()

	// Handler
	handler := userhandler.New(app, validator, userCase)

	usersGroup := group.Group("/users")
	{
		publicRouter := usersGroup.Group("")
		{
			publicRouter.POST("/refresh-token", handler.RefreshToken)
		}

		authRouter := usersGroup.Group("/auth")
		{
			authRouter.POST("/register", handler.Register)
			authRouter.POST("/login", handler.Login)
		}

		protectedRouter := usersGroup.Group("")
		protectedRouter.Use(middleware.Auth(app.Config.Auth, authCase))
		{
			protectedRouter.GET("/profile", handler.Profile)
			protectedRouter.POST("/:id/tasks", handler.CreateTask, middleware.CheckIsValidUserID)
			protectedRouter.GET("/:id/tasks", handler.Tasks)
		}
	}
}
