package taskvalidator

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go-backend-clean-arch/internal/domain/dto/taskdto"
	"go-backend-clean-arch/internal/infrastructure/kind"
	"go-backend-clean-arch/internal/infrastructure/richerror"
	"go-backend-clean-arch/pkg/message"
)

func (v Validator) ValidateCreateRequest(req taskdto.CreateRequest) (map[string]string, error) {
	const op = message.OpTaskValidatorValidateCreateRequest

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Title,
			validation.Required,
			validation.Length(3, 128)),

		validation.Field(&req.Description,
			validation.Required,
			validation.Length(12, 1024)),
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
