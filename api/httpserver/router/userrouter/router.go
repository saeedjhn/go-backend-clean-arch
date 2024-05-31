package userrouter

import (
	"github.com/labstack/echo/v4"
	"go-backend-clean-arch-according-to-go-standards-project-layout/api/httpserver/handler/userhandler"
	"go-backend-clean-arch-according-to-go-standards-project-layout/api/httpserver/intercaptor/userinterceptor"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/bootstrap"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/gateway/taskinggateway"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/repository/taskrepository/postgresqltask"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/repository/userrespository/postgresqluser"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/usecase/taskusecase"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/usecase/userusecase"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/validator/uservalidator"
)

func New(
	app *bootstrap.Application,
	e *echo.Group,
) {
	// Repository
	tdb := postgresqltask.New(app.PostgresqlDB)
	udb := postgresqluser.New(app.PostgresqlDB)

	// Repository & Usecase
	tu := taskusecase.New(tdb)

	// Service-oriented - no depends on useCases - ( userusecase -> taskgateway -> taskusecase )
	tg := taskinggateway.New(tu)

	// Repository & Usecase
	uu := userusecase.New(tg, udb)

	// Validator
	uv := uservalidator.New()

	// Handler
	userHandler := userhandler.New(app, uv, uu)

	g := e.Group("/users")

	g.Use(userinterceptor.TransformResponse(app.Config.Application.Env))

	publicRouter := g.Group("/auth")
	{
		publicRouter.POST("/register", userHandler.Register)
		publicRouter.POST("/login", userHandler.Login)
	}

	protectedRouter := g.Group("")

	//protectedRouter.Use(middleware.Auth())
	{
		protectedRouter.GET("/task-list", userHandler.TaskList)
	}
}
