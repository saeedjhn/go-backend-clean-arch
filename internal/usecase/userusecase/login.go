package userusecase

import (
	"context"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/servicedto/userauthservicedto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/userdto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/kind"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/message"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/security/bcrypt"
)

func (i *Interactor) Login(ctx context.Context, req userdto.LoginRequest) (userdto.LoginResponse, error) {
	user, err := i.repository.GetByMobile(ctx, req.Mobile)
	if err != nil {
		return userdto.LoginResponse{}, err
	}

	err = bcrypt.CompareHashAndSTR(user.Password, req.Password)
	if err != nil {
		return userdto.LoginResponse{}, richerror.New(_opUserServiceLogin).WithErr(err).
			WithMessage(_errMsgIncorrectPassword).
			WithKind(kind.KindStatusBadRequest)
	}

	dto := userauthservicedto.CreateTokenRequest{User: entity.User{
		ID:        user.ID,
		Name:      user.Name,
		Mobile:    user.Mobile,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}}

	accessToken, err := i.authIntr.CreateAccessToken(dto)
	if err != nil {
		return userdto.LoginResponse{}, richerror.New(_opUserServiceLogin).WithErr(err).
			WithMessage(message.ErrorMsg500InternalServerError).
			WithKind(kind.KindStatusInternalServerError)
	}

	refreshToken, err := i.authIntr.CreateRefreshToken(dto)
	if err != nil {
		return userdto.LoginResponse{}, richerror.New(_opUserServiceLogin).WithErr(err).
			WithMessage(message.ErrorMsg500InternalServerError).
			WithKind(kind.KindStatusInternalServerError)
	}

	return userdto.LoginResponse{
		User: userdto.UserInfo{
			ID:        user.ID,
			Name:      user.Name,
			Mobile:    user.Mobile,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}, // Or
		//User: user.ToUserInfoDTO(),
		Token: userdto.Token{
			AccessToken:  accessToken.Token,
			RefreshToken: refreshToken.Token,
		},
	}, nil
}
