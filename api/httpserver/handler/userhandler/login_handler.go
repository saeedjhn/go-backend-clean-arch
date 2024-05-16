package userhandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (u *UserHandler) Login(c echo.Context) error {
	//req := userdto.UserRequest{}

	// Call usecase

	return c.JSON(http.StatusOK, "UserHandler -> Login - IMPL ME")
}
