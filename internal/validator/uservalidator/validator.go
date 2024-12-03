package uservalidator

import (
	"github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/handler/userhandler"
	"github.com/saeedjhn/go-backend-clean-arch/configs"
)

type Validator struct {
	config *configs.Config
}

var _ userhandler.Validator = (*Validator)(nil)

func New(config *configs.Config) Validator {
	return Validator{config: config}
}
