package uservalidator

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go-backend-clean-arch/internal/dto/userdto"
	"go-backend-clean-arch/internal/infrastructure/kind"
	"go-backend-clean-arch/internal/infrastructure/richerror"
	"go-backend-clean-arch/pkg/message"
)

func (v Validator) ValidateProfileRequest(req userdto.ProfileRequest) (map[string]string, error) {
	const op = message.OpUserValidatorValidateLoginRequest

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.ID,
			validation.Required),
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
