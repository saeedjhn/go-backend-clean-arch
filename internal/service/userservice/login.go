package userservice

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/servicedto/userauthservicedto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/userdto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/kind"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/richerror"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/security/bcrypt"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/message"
)

func (u *UserInteractor) Login(req userdto.LoginRequest) (userdto.LoginResponse, error) {
	user, err := u.repository.GetByMobile(req.Mobile)
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

	accessToken, err := u.authInteractor.CreateAccessToken(dto)
	if err != nil {
		return userdto.LoginResponse{}, richerror.New(_opUserServiceLogin).WithErr(err).
			WithMessage(message.ErrorMsg500InternalServerError).
			WithKind(kind.KindStatusInternalServerError)
	}

	refreshToken, err := u.authInteractor.CreateRefreshToken(dto)
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
