package userrouter

import (
	"github.com/labstack/echo/v4"
	"go-backend-clean-arch-according-to-go-standards-project-layout/api/httpserver/handler/userhandler"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/bootstrap"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/gateway/taskgateway"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/gateway/usergateway"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/repository/taskrepository/postgresqltask"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/repository/userrespository/postgresqluser"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/usecase/taskusecase"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/usecase/userusecase"
)

func New(
	app *bootstrap.Application,
	e *echo.Echo,
) {
	// Repository
	taskPostgresql := postgresqltask.New(app.PostgresqlDB)
	userPostgresql := postgresqluser.New(app.PostgresqlDB)

	// Gateway & Usecase
	taskGateway := taskgateway.New(taskPostgresql)
	taskUsecase := taskusecase.New(taskGateway)

	// Gateway & Usecase
	userGateway := usergateway.New(userPostgresql, taskUsecase)
	userUsecase := userusecase.New(userGateway)

	// Handler
	userHandler := userhandler.New(app, userUsecase)

	g := e.Group("/users")

	publicRouter := g.Group("/auth")
	{
		publicRouter.POST("/register", userHandler.Register)
		publicRouter.POST("/login", userHandler.Login)
	}

	protectedRouter := publicRouter.Group("")

	//protectedRouter.Use(middleware.Auth())
	{
		protectedRouter.GET("/task-list", userHandler.TaskList)
	}
}
