package userusecase

import (
	"context"
	"errors"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/kind"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/security/bcrypt"

	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/userdto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
)

func (i *Interactor) Register(ctx context.Context, req userdto.RegisterRequest) (userdto.RegisterResponse, error) {
	isUnique, err := i.repository.IsMobileUnique(ctx, req.Mobile)
	if err != nil {
		return userdto.RegisterResponse{}, err
	}

	if !isUnique {
		return userdto.RegisterResponse{},
			richerror.New(_opUserServiceRegister).
				WithErr(errors.New(_errMsgMobileIsNotUnique)).
				WithMessage(_errMsgMobileIsNotUnique).
				WithKind(kind.KindStatusBadRequest)
	}

	user := entity.User{
		Name:   req.Name,
		Mobile: req.Mobile,
	}

	encryptedPass, _ := bcrypt.Generate(req.Password, bcrypt.DefaultCost) // Check err
	user.Password = encryptedPass

	createdUser, err := i.repository.Create(ctx, user)
	if err != nil {
		return userdto.RegisterResponse{}, err
	}

	return userdto.RegisterResponse{
		User: userdto.UserInfo{
			ID:        createdUser.ID,
			Name:      createdUser.Name,
			Mobile:    createdUser.Mobile,
			Email:     createdUser.Email,
			CreatedAt: createdUser.CreatedAt,
			UpdatedAt: createdUser.UpdatedAt,
		}, // Or
		//User: createdUser.ToUserInfoDTO(),
	}, nil
}
