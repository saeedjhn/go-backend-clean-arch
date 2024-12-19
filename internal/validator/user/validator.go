package user

import (
	"github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/handler/user"
	"github.com/saeedjhn/go-backend-clean-arch/configs"
)

type Validator struct {
	config *configs.Config
}

var _ user.Validator = (*Validator)(nil)

func New(config *configs.Config) Validator {
	return Validator{config: config}
}
