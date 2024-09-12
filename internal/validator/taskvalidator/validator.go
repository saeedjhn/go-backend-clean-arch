package taskvalidator

import "github.com/saeedjhn/go-backend-clean-arch/api/httpserver/handler/taskhandler"

type Validator struct {
}

var _ taskhandler.Validator = (*Validator)(nil)

func New() Validator {
	return Validator{}
}
