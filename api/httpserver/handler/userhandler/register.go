package userhandler

import (
	"github.com/labstack/echo/v4"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/dto/userdto"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/bind"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/httpstatus"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/response/httpresponse"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/richerror"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/sanitize"
	"go-backend-clean-arch-according-to-go-standards-project-layout/pkg/message"
	"net/http"
)

func (u *UserHandler) Register(c echo.Context) error {
	// Initial
	req := userdto.RegisterRequest{}
	httpRes := httpresponse.New()

	// Bind
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,
			httpRes.
				WithStatusCode(http.StatusBadRequest).
				WithMessage(message.ErrorMsg400BadRequest).
				WithError(bind.CheckErrFromBind(err)).Build(),
		)
	}

	// Validation
	if fieldsErrs, err := u.userValidator.ValidateRegisterRequest(req); err != nil {
		richErr, _ := richerror.Analysis(err)
		code := httpstatus.FromKind(richErr.Kind())

		return c.JSON(
			code,
			httpRes.
				WithStatusCode(code).
				WithMessage(richErr.Message()).
				WithError(fieldsErrs).Build(),
		)
	}

	// Sanitize
	sanitize.New().
		SetPolicy(sanitize.StrictPolicy).
		Struct(&req)

	// UseCase
	uRes, err := u.userInteractor.Register(req)
	if err != nil {
		richErr, _ := richerror.Analysis(err)
		code := httpstatus.FromKind(richErr.Kind())

		return c.JSON(
			code,
			httpRes.
				WithStatusCode(code).
				WithMessage(richErr.Message()).
				WithError(richErr.Error()).Build(),
		)
	}

	return c.JSON(
		http.StatusCreated,
		httpRes.
			WithStatus(true).
			WithStatusCode(http.StatusCreated).
			WithMessage(message.MsgUserRegisterSuccessfully).
			WithMeta(uRes).Build(),
	)
}
