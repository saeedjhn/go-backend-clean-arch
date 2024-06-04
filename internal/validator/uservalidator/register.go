package uservalidator

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/dto/userdto"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/kind"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/richerror"
	"go-backend-clean-arch-according-to-go-standards-project-layout/pkg/message"
)

func (v Validator) ValidateRegisterRequest(req userdto.RegisterRequest) (map[string]string, error) {
	const op = message.OpUserValidatorValidateRegisterRequest

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Name,
			validation.Required,
			validation.Length(3, 128)),

		validation.Field(&req.Mobile,
			validation.Required,
			validation.Length(11, 11)),
	); err != nil {
		fieldErrors := make(map[string]string)

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

	return nil, nil

}
