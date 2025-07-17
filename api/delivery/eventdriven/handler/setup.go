package handler

import (
	usereventhandler "github.com/saeedjhn/go-backend-clean-arch/api/delivery/eventdriven/handler/user"
	"github.com/saeedjhn/go-backend-clean-arch/api/delivery/eventdriven/registery"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
)

func Setup(app *bootstrap.Application) types.EventRouter {
	er := registery.New()

	usereventhandler.Setup(app, er)

	return er.Handlers()
}
