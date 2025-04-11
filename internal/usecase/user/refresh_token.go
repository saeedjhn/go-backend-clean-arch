package user

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"

	"github.com/saeedjhn/go-backend-clean-arch/internal/dto/user"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/msg"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"
)

func (i *Interactor) RefreshToken(
	ctx context.Context,
	req user.RefreshTokenRequest,
) (user.RefreshTokenResponse, error) {
	ctx, span := i.trc.Span(ctx, "RefreshToken")
	span.SetAttributes(map[string]interface{}{
		"usecase.name":    "RefreshToken",
		"usecase.request": req,
	})
	defer span.End()

	if fieldsErrs, err := i.vld.ValidateRefreshTokenRequest(req); err != nil {
		return user.RefreshTokenResponse{FieldErrors: fieldsErrs}, err
	}

	resp, err := i.authIntr.ParseToken(i.cfg.Auth.RefreshTokenSecret, req.RefreshToken)
	if err != nil {
		return user.RefreshTokenResponse{}, richerror.New(_opUserServiceRefreshToken).WithErr(err).
			WithMessage(msg.ErrMsg403Forbidden).
			WithKind(richerror.KindStatusBadRequest)
	}

	u, err := i.repository.GetByID(ctx, resp.UserID.Uint64())
	if err != nil {
		return user.RefreshTokenResponse{}, err
	}

	authenticable := models.Authenticable{ID: u.ID}

	accessToken, err := i.authIntr.CreateAccessToken(authenticable)
	if err != nil {
		return user.RefreshTokenResponse{}, richerror.New(_opUserServiceRefreshToken).WithErr(err).
			WithMessage(msg.ErrMsg400BadRequest).
			WithKind(richerror.KindStatusBadRequest)
	}

	refreshToken, err := i.authIntr.CreateRefreshToken(authenticable)
	if err != nil {
		return user.RefreshTokenResponse{}, richerror.New(_opUserServiceRefreshToken).WithErr(err).
			WithMessage(msg.ErrMsg400BadRequest).
			WithKind(richerror.KindStatusBadRequest)
	}

	return user.RefreshTokenResponse{
		Tokens: user.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}
