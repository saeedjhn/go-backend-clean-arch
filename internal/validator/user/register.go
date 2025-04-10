package user

import (
	"errors"

	"github.com/go-ozzo/ozzo-validation/v4/is"

	userdto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/user"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	passwordvalidator "github.com/wagslane/go-password-validator"
)

func (v Validator) ValidateRegisterRequest(req userdto.RegisterRequest) (map[string]string, error) {
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Name,
			validation.Required,
			validation.Length(_nameMinLen, _nameMaxLen)),

		validation.Field(&req.Email,
			validation.Required,
			is.Email),

		validation.Field(&req.Mobile,
			validation.Required,
			validation.Length(_mobileMinLen, _mobileMaxLen)),

		validation.Field(&req.Password,
			validation.Required,
			validation.Length(_passMinLen, _passMaxLen),
			validation.By(isSecurePassword(v.entropyPassword))),
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

		return fieldErrors, richerror.New(_opUserValidatorValidateRegisterRequest).WithErr(err).
			WithMessage(errMsgInvalidInput).
			WithKind(richerror.KindStatusUnprocessableEntity)
	}

	return nil, nil //nolint:nilnil // return both the `nil` error and invalid value: use a sentinel error instead (nilnil)
}

func isSecurePassword(entropy float64) func(value interface{}) error {
	return func(value interface{}) error {
		p, _ := value.(string)

		if err := passwordvalidator.Validate(p, entropy); err != nil {
			return errors.New(
				"insecure password, try including more special characters, " +
					"using uppercase letters, using numbers or using a longer password",
			)
		}

		return nil
	}
}
