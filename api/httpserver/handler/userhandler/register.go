package userhandler

import (
	"github.com/labstack/echo/v4"
	"go-backend-clean-arch/internal/domain/dto/userdto"
	"go-backend-clean-arch/internal/infrastructure/bind"
	"go-backend-clean-arch/internal/infrastructure/httpstatus"
	"go-backend-clean-arch/internal/infrastructure/richerror"
	"go-backend-clean-arch/internal/infrastructure/sanitize"
	"go-backend-clean-arch/pkg/message"
	"net/http"
)

func (u *UserHandler) Register(c echo.Context) error {
	// Bind
	var req = userdto.RegisterRequest{}
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

		return echo.NewHTTPError(code,
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
		Struct(&req) // Check err

	// Usage Use-case
	resp, err := u.userInteractor.Register(req)
	if err != nil {
		richErr, _ := richerror.Analysis(err)
		code := httpstatus.FromKind(richErr.Kind())

		return echo.NewHTTPError(code,
			echo.Map{
				"status":  false,
				"message": richErr.Message(),
				"errors":  richErr.Error(),
			})
	}
	return c.JSON(http.StatusCreated,
		echo.Map{
			"status":  true,
			"message": message.MsgUserRegisterSuccessfully,
			"data":    resp,
		},
	)
}
