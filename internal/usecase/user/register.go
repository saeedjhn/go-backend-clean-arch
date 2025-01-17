package user

import (
	"context"
	"errors"

	"github.com/saeedjhn/go-backend-clean-arch/internal/dto/user"
	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"
)

func (i *Interactor) Register(ctx context.Context, req user.RegisterRequest) (user.RegisterResponse, error) {
	ctx, span := i.trc.Span(ctx, "Register")
	span.SetAttributes(map[string]interface{}{
		"usecase.name":    "Register",
		"usecase.request": req,
	})
	defer span.End()

	if fieldsErrs, err := i.vld.ValidateRegisterRequest(req); err != nil {
		return user.RegisterResponse{FieldErrors: fieldsErrs}, err
	}

	isUnique, err := i.repository.IsMobileUnique(ctx, req.Mobile)
	if err != nil {
		return user.RegisterResponse{}, err
	}

	if !isUnique {
		return user.RegisterResponse{},
			richerror.New(_opUserServiceRegister).
				WithErr(errors.New(_errMsgMobileIsNotUnique)).
				WithMessage(_errMsgMobileIsNotUnique).
				WithKind(richerror.KindStatusBadRequest)
	}

	u := entity.User{
		Name:   req.Name,
		Mobile: req.Mobile,
	}

	encryptedPass, _ := GenerateHash(req.Password) // TODO: usecase>userusecase>register>CheckErrorGenerateHash
	u.Password = encryptedPass

	createdUser, err := i.repository.Create(ctx, u)
	if err != nil {
		return user.RegisterResponse{}, err
	}

	return user.RegisterResponse{
		Data: user.Data{
			ID:        createdUser.ID,
			Name:      createdUser.Name,
			Mobile:    createdUser.Mobile,
			Email:     createdUser.Email,
			CreatedAt: createdUser.CreatedAt,
			UpdatedAt: createdUser.UpdatedAt,
		}, // Or
		// Data: createdUser.ToUserInfoDTO(),
	}, nil
}
