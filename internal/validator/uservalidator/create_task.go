package uservalidator

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/userdto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/kind"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/richerror"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/message"
)

func (v Validator) ValidateCreateTaskRequest(req userdto.CreateTaskRequest) (map[string]string, error) {
	const op = message.OpTaskValidatorValidateCreateTaskRequest

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Title,
			validation.Required,
			validation.Length(TitleMinLen, TitleMaxLen)),

		validation.Field(&req.Description,
			validation.Required,
			validation.Length(DescMinLen, DescMaxLen)),
	); err != nil {
		var fieldErrors = make(map[string]string)

		var errV validation.Errors
		ok := errors.As(err, &errV)
		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}

		return fieldErrors, richerror.New(op).WithErr(err).
			WithMessage(message.ErrorMsgInvalidInput).
			WithKind(kind.KindStatusUnprocessableEntity)
	}

	return map[string]string{}, nil
}
