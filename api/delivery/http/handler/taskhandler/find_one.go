package taskhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) FindOne(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"status":  true,
		"message": "FIND ONE",
		"data":    "",
	})
}
