package userusecase

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/userdto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/kind"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/message"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/security/bcrypt"
)

func (i *Interactor) Login(ctx context.Context, req userdto.LoginRequest) (userdto.LoginResponse, error) {
	ctx, span := i.trc.Span(ctx, "Login")
	span.SetAttributes(map[string]interface{}{
		"usecase.name":    "Login",
		"usecase.request": req,
	})
	defer span.End()

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

	authenticable := entity.Authenticable{ID: user.ID}

	accessToken, err := i.authIntr.CreateAccessToken(authenticable)
	if err != nil {
		return userdto.LoginResponse{}, richerror.New(_opUserServiceLogin).WithErr(err).
			WithMessage(message.ErrorMsg500InternalServerError).
			WithKind(kind.KindStatusInternalServerError)
	}

	refreshToken, err := i.authIntr.CreateRefreshToken(authenticable)
	if err != nil {
		return userdto.LoginResponse{}, richerror.New(_opUserServiceLogin).WithErr(err).
			WithMessage(message.ErrorMsg500InternalServerError).
			WithKind(kind.KindStatusInternalServerError)
	}

	return userdto.LoginResponse{
		Data: userdto.UserInfo{
			ID:        user.ID,
			Name:      user.Name,
			Mobile:    user.Mobile,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}, // Or
		// Data: user.ToUserInfoDTO(),
		Tokens: userdto.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}
