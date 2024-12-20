package user

import (
	"errors"

	"github.com/saeedjhn/go-backend-clean-arch/internal/dto/user"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/kind"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateRefreshTokenRequest(req user.RefreshTokenRequest) (map[string]string, error) {
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
