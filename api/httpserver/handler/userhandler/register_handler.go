package userhandler

import (
	"github.com/labstack/echo/v4"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/dto/userdto"
	"net/http"
)

func (u *UserHandler) Register(c echo.Context) error {
	req := userdto.UserRequest{}

	u.userInteractor.Register(req)
	return c.JSON(http.StatusOK, "UserHandler -> Register - IMPL ME")
}
