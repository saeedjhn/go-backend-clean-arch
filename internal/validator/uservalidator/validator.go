package uservalidator

import (
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/delivery/http/handler/userhandler"
)

type Validator struct {
	config *configs.Config
}

var _ userhandler.Validator = (*Validator)(nil)

func New(config *configs.Config) Validator {
	return Validator{config: config}
}
