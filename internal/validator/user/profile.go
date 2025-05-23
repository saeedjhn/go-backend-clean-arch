package user

import (
	"errors"

	userdto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/user"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateProfileRequest(req userdto.ProfileRequest) (map[string]string, error) {
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.ID,
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

		return fieldErrors, richerror.New(_opUserValidatorValidateProfileRequest).WithErr(err).
			WithMessage(errMsgInvalidInput).
			WithKind(richerror.KindStatusUnprocessableEntity)
	}

	return nil, nil //nolint:nilnil // return both the `nil` error and invalid value: use a sentinel error instead (nilnil)
}
