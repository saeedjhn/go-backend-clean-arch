package uservalidator

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/userdto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/kind"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/richerror"
)

func (v Validator) ValidateRefreshTokenRequest(req userdto.RefreshTokenRequest) (map[string]string, error) {
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.RefreshToken,
			validation.Required),
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

		return fieldErrors, richerror.New(_opUserValidatorValidateRefTokenRequest).WithErr(err).
			WithMessage(_errMsgInvalidInput).
			WithKind(kind.KindStatusUnprocessableEntity)
	}

	return map[string]string{}, nil
}
