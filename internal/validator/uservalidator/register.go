package uservalidator

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/userdto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/kind"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/richerror"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/message"
	passwordvalidator "github.com/wagslane/go-password-validator"
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

		validation.Field(&req.Password,
			validation.Required,
			validation.Length(3, 128),
			validation.By(isSecurePassword(v.config.Application.EntropyPassword))),
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

func isSecurePassword(entropy float64) func(value interface{}) error {
	return func(value interface{}) error {
		p, _ := value.(string)

		if err := passwordvalidator.Validate(p, entropy); err != nil {
			return errors.New(
				"insecure password, try including more special characters, using uppercase letters, using numbers or using a longer password",
			)
		}

		return nil
	}
}
