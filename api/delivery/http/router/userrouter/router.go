package userrouter

import (
	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/handler/userhandler"
	"github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/middleware"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"github.com/saeedjhn/go-backend-clean-arch/internal/validator/uservalidator"
)

func New(
	app *bootstrap.Application,
	group *echo.Group,
) {
	// Validator
	validator := uservalidator.New(app.Config)

	// Handler
	handler := userhandler.New(
		app,
		app.Trc,
		validator,
		app.Usecase.UserIntr,
		app.Usecase.TaskIntr,
	)

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
		protectedRouter.Use(middleware.Auth(app.Usecase.AuthIntr))
		{
			protectedRouter.GET("/profile", handler.Profile)
			protectedRouter.POST("/:id/tasks", handler.CreateTask, middleware.CheckIsValidUserID)
			protectedRouter.GET("/:id/tasks", handler.Tasks, middleware.CheckIsValidUserID)
		}
	}
}
