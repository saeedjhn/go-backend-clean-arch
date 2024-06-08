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

func (u *UserHandler) Tasks(c echo.Context) error {
	// Initial
	req := userdto.TasksRequest{}

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

	// Usage Use-case
	resp, err := u.userInteractor.Tasks(req)
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
			"message": message.MsgUserGetTasksSuccessfully,
			"data":    resp,
		},
	)
}
