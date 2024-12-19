package user

import (
	"context"
	"github.com/saeedjhn/go-backend-clean-arch/internal/dto/user"
	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/kind"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/message"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/security/bcrypt"
)

func (i *Interactor) Login(ctx context.Context, req user.LoginRequest) (user.LoginResponse, error) {
	ctx, span := i.trc.Span(ctx, "Login")
	span.SetAttributes(map[string]interface{}{
		"usecase.name":    "Login",
		"usecase.request": req,
	})
	defer span.End()

	u, err := i.repository.GetByMobile(ctx, req.Mobile)
	if err != nil {
		return user.LoginResponse{}, err
	}

	err = bcrypt.CompareHashAndSTR(u.Password, req.Password)
	if err != nil {
		return user.LoginResponse{}, richerror.New(_opUserServiceLogin).WithErr(err).
			WithMessage(_errMsgIncorrectPassword).
			WithKind(kind.KindStatusBadRequest)
	}

	authenticable := entity.Authenticable{ID: u.ID}

	accessToken, err := i.authIntr.CreateAccessToken(authenticable)
	if err != nil {
		return user.LoginResponse{}, richerror.New(_opUserServiceLogin).WithErr(err).
			WithMessage(message.ErrorMsg500InternalServerError).
			WithKind(kind.KindStatusInternalServerError)
	}

	refreshToken, err := i.authIntr.CreateRefreshToken(authenticable)
	if err != nil {
		return user.LoginResponse{}, richerror.New(_opUserServiceLogin).WithErr(err).
			WithMessage(message.ErrorMsg500InternalServerError).
			WithKind(kind.KindStatusInternalServerError)
	}

	return user.LoginResponse{
		Data: user.UserInfo{
			ID:        u.ID,
			Name:      u.Name,
			Mobile:    u.Mobile,
			Email:     u.Email,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		}, // Or
		// Data: user.ToUserInfoDTO(),
		Tokens: user.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}
