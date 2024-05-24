package router

import (
	"github.com/labstack/echo/v4"
	"go-backend-clean-arch-according-to-go-standards-project-layout/api/httpserver/router/healthcheckrouter"
	"go-backend-clean-arch-according-to-go-standards-project-layout/api/httpserver/router/taskrouter"
	"go-backend-clean-arch-according-to-go-standards-project-layout/api/httpserver/router/userrouter"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/bootstrap"
)

func Setup(
	app *bootstrap.Application,
	e *echo.Echo,
) {
	g := e.Group("")

	userrouter.New(app, g)
	taskrouter.New(app, g)
	healthcheckrouter.New(app, g)
}
