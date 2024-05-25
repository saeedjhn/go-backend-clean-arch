package userhandler

import (
	"github.com/labstack/echo/v4"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/dto/userdto"
	"net/http"
)

func (u *UserHandler) TaskList(c echo.Context) error {
	req := userdto.TaskListRequest{}

	u.userInteractor.TaskList(req)

	return c.JSON(http.StatusOK, "UserHandler -> TaskList - IMPL ME")
}
