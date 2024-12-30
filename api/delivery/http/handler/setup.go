package handler

import (
	"github.com/labstack/echo/v4"
	adminhandler "github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/handler/admin"
	healthcheckhandler "github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/handler/healthcheck"
	prometheushandler "github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/handler/prometheus"
	taskhandler "github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/handler/user/task"
	userhandler "github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/handler/user/user"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
)

func Setup(app *bootstrap.Application, e *echo.Echo) {
	adminhandler.Setup(app, e)
	userhandler.Setup(app, e)
	taskhandler.Setup(app, e)
	prometheushandler.Setup(app, e)
	healthcheckhandler.Setup(app, e)
}
