package taskhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Create(c echo.Context) error {
	return c.JSON(http.StatusCreated, echo.Map{
		"status":  true,
		"message": "CREATE",
		"data":    "",
	})
}
