package prometheus

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
)

func New(_ *bootstrap.Application, group *echo.Group) {
	// adds route to serve gathered metrics
	group.GET("/metrics", echoprometheus.NewHandler())
}
