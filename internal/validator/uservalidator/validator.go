package uservalidator

import "github.com/saeedjhn/go-backend-clean-arch/api/httpserver/handler/userhandler"

type Validator struct {
}

var _ userhandler.Validator = (*Validator)(nil)

func New() Validator {
	return Validator{}
}
