package handler

import (
	"github.com/labstack/echo/v4"
	adminhandler "github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/handler/admin"
	csrfhandler "github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/handler/csrf"
	healthcheckhandler "github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/handler/healthz"
	prometheushandler "github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/handler/prometheus"
	userhandler "github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/handler/user/user"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
)

func Setup(app *bootstrap.Application, e *echo.Echo) {
	adminhandler.Setup(app, e)
	userhandler.Setup(app, e)
	prometheushandler.Setup(app, e)
	csrfhandler.Setup(app, e)
	healthcheckhandler.Setup(app, e)
}
