package user

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/internal/dto/user"
	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/msg"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"
)

func (i *Interactor) Login(ctx context.Context, req user.LoginRequest) (user.LoginResponse, error) {
	ctx, span := i.trc.Span(ctx, "Login")
	span.SetAttributes(map[string]interface{}{
		"usecase.name":    "Login",
		"usecase.request": req,
	})
	defer span.End()

	if fieldsErrs, err := i.vld.ValidateLoginRequest(req); err != nil {
		return user.LoginResponse{FieldErrors: fieldsErrs}, err
	}

	u, err := i.repository.GetByMobile(ctx, req.Mobile)
	if err != nil {
		return user.LoginResponse{}, err
	}

	err = CompareHash(u.Password, req.Password)
	if err != nil {
		return user.LoginResponse{}, richerror.New(_opUserServiceLogin).WithErr(err).
			WithMessage(errMsgIncorrectPassword).
			WithKind(richerror.KindStatusBadRequest)
	}

	authenticable := entity.Authenticable{ID: u.ID}

	accessToken, err := i.authIntr.CreateAccessToken(authenticable)
	if err != nil {
		return user.LoginResponse{}, richerror.New(_opUserServiceLogin).WithErr(err).
			WithMessage(msg.ErrorMsg500InternalServerError).
			WithKind(richerror.KindStatusInternalServerError)
	}

	refreshToken, err := i.authIntr.CreateRefreshToken(authenticable)
	if err != nil {
		return user.LoginResponse{}, richerror.New(_opUserServiceLogin).WithErr(err).
			WithMessage(msg.ErrorMsg500InternalServerError).
			WithKind(richerror.KindStatusInternalServerError)
	}

	return user.LoginResponse{
		UserInfo: user.UserInfo{
			ID:        u.ID,
			Name:      u.Name,
			Mobile:    u.Mobile,
			Email:     u.Email,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		},
		Tokens: user.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}
