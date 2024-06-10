package userhandler

import (
	"github.com/labstack/echo/v4"
	"go-backend-clean-arch/internal/dto/userdto"
	"go-backend-clean-arch/internal/infrastructure/bind"
	"go-backend-clean-arch/internal/infrastructure/httpstatus"
	"go-backend-clean-arch/internal/infrastructure/richerror"
	"go-backend-clean-arch/pkg/message"
	"net/http"
)

func (u *UserHandler) CreateTask(c echo.Context) error {
	// Initial
	req := userdto.CreateTaskRequest{}

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
	if fieldsErrs, err := u.userValidator.ValidateCreateTaskRequest(req); err != nil {
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
	//sanitize.New().
	//	SetPolicy(sanitize.StrictPolicy).
	//	Struct(&req) // nolint:errcheck

	// Usage Use-case
	resp, err := u.userInteractor.CreateTask(req)
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
		http.StatusOK,
		echo.Map{
			"status":  true,
			"message": message.MsgUserCreateTaskSuccessfully,
			"data":    resp,
		},
	)
}
