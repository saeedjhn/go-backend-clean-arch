package userrouter

import (
	"github.com/labstack/echo/v4"
	"go-backend-clean-arch/api/httpserver/handler/userhandler"
	"go-backend-clean-arch/api/httpserver/middleware"
	"go-backend-clean-arch/internal/bootstrap"
	"go-backend-clean-arch/internal/gateway/usergateway"
	"go-backend-clean-arch/internal/infrastructure/token"
	"go-backend-clean-arch/internal/repository/taskrepository/mysqltask"
	"go-backend-clean-arch/internal/repository/userrespository/mysqluser"
	"go-backend-clean-arch/internal/usecase/authusecase"
	"go-backend-clean-arch/internal/usecase/taskusecase"
	"go-backend-clean-arch/internal/usecase/userusecase"
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
	taskCase := taskusecase.New(taskMysql)
	authCase := authusecase.New(app.Config.Auth, token.New())

	// Service-oriented - no depends on use-cases - ( userusecase -> usergateway -> taskusecase )
	userGate := usergateway.New(authCase, taskCase)

	// Repository & Usecase
	userCase := userusecase.New(app.Config, userGate, userMysql)

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
	protectedRouter.Use(middleware.Auth(app.Config.Auth, authCase))
	{
		protectedRouter.GET("/profile", handler.Profile)
		//protectedRouter.GET("/tasks", handler.TaskList)
		protectedRouter.GET("/:id/tasks", handler.Tasks)
	}
}
