package router

import (
	"github.com/labstack/echo/v4"
	"go-backend-clean-arch/api/httpserver/router/healthcheckrouter"
	"go-backend-clean-arch/api/httpserver/router/taskrouter"
	"go-backend-clean-arch/api/httpserver/router/userrouter"
	"go-backend-clean-arch/internal/bootstrap"
)

func Setup(
	app *bootstrap.Application,
	echo *echo.Echo,
) {
	routerGroup := echo.Group("")

	userrouter.New(app, routerGroup)
	taskrouter.New(app, routerGroup)
	healthcheckrouter.New(app, routerGroup)
}
