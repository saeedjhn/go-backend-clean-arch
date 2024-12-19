package router

import (
	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/router/healthcheck"
	"github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/router/prometheus"
	"github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/router/task"
	"github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/router/user"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
)

func Setup(
	app *bootstrap.Application,
	echo *echo.Echo,
) {
	routerGroup := echo.Group("") // Parent

	user.New(app, routerGroup)
	task.New(app, routerGroup)
	prometheus.New(app, routerGroup)
	healthcheck.New(app, routerGroup)
}
