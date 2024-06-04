package userhandler

import (
	"github.com/labstack/echo/v4"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/dto/userdto"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/bind"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/httpstatus"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/richerror"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/sanitize"
	"go-backend-clean-arch-according-to-go-standards-project-layout/pkg/message"
	"net/http"
)

func (u *UserHandler) Register(c echo.Context) error {
	// Initial
	req := userdto.RegisterRequest{}

	// Bind
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,
			echo.Map{
				"status":  false,
				"message": message.ErrorMsg400BadRequest,
				"errors":  bind.CheckErrFromBind(err).Error(),
			},
		)
	}

	// Validation
	if fieldsErrs, err := u.userValidator.ValidateRegisterRequest(req); err != nil {
		richErr, _ := richerror.Analysis(err)
		code := httpstatus.FromKind(richErr.Kind())

		return echo.NewHTTPError(
			code,
			echo.Map{
				"status":  false,
				"message": richErr.Message(),
				"errors":  fieldsErrs,
			},
		)
	}

	// Sanitize
	sanitize.New().
		SetPolicy(sanitize.StrictPolicy).
		Struct(&req)

	// Usage Use-case
	resp, err := u.userInteractor.Register(req)
	if err != nil {
		richErr, _ := richerror.Analysis(err)
		code := httpstatus.FromKind(richErr.Kind())

		return echo.NewHTTPError(
			code,
			echo.Map{
				"status":  false,
				"message": richErr.Message(),
				"errors":  richErr.Error(),
			})

	}

	return c.JSON(
		http.StatusCreated,
		echo.Map{
			"status":  true,
			"message": message.MsgUserRegisterSuccessfully,
			"data":    resp,
		},
	)
}
