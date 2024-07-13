package userrouter

import (
	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/api/httpserver/handler/userhandler"
	"github.com/saeedjhn/go-backend-clean-arch/api/httpserver/middleware"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/token"
	"github.com/saeedjhn/go-backend-clean-arch/internal/repository/taskrepository/mysqltask"
	"github.com/saeedjhn/go-backend-clean-arch/internal/repository/userrespository/mysqluser"
	"github.com/saeedjhn/go-backend-clean-arch/internal/service/authservice"
	"github.com/saeedjhn/go-backend-clean-arch/internal/service/taskservice"
	"github.com/saeedjhn/go-backend-clean-arch/internal/service/userservice"
	"github.com/saeedjhn/go-backend-clean-arch/internal/validator/uservalidator"
)

func New(
	app *bootstrap.Application,
	group *echo.Group,
) {
	// Repository
	taskMysql := mysqltask.New(app.MysqlDB)
	userMysql := mysqluser.New(app.MysqlDB)

	// Usecase
	taskSvc := taskservice.New(taskMysql)
	authSvc := authservice.New(app.Config.Auth, token.New())

	// Service-oriented - inject service to another service
	// Repository & Usecase
	userSvc := userservice.New(
		app.Config, authSvc, taskSvc, userMysql,
	)

	// Validator
	validator := uservalidator.New()

	// Handler
	handler := userhandler.New(app, validator, userSvc)

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
		protectedRouter.Use(middleware.Auth(app.Config.Auth, authSvc))
		{
			protectedRouter.GET("/profile", handler.Profile)
			protectedRouter.POST("/:id/tasks", handler.CreateTask, middleware.CheckIsValidUserID)
			protectedRouter.GET("/:id/tasks", handler.Tasks)
		}
	}
}
