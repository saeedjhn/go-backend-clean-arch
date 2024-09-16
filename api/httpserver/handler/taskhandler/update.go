package taskhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (t *TaskHandler) Update(c echo.Context) error {

	return c.JSON(http.StatusOK, echo.Map{
		"status":  true,
		"message": "UPDATE",
		"data":    "",
	})
}
