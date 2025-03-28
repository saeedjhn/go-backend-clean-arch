package task

import (
	"errors"

	taskdto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/task"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateCreateRequest(req taskdto.CreateRequest) (map[string]string, error) {
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Title,
			validation.Required,
			validation.Length(_titleMinLen, _titleMaxLen)),

		validation.Field(&req.Description,
			validation.Required,
			validation.Length(_descMinLen, _descMaxLen)),
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

		return fieldErrors, richerror.New(_opTaskValidatorValidateCreateRequest).WithErr(err).
			WithMessage(errMsgInvalidInput).
			WithKind(richerror.KindStatusUnprocessableEntity)
	}

	return map[string]string{}, nil
}
