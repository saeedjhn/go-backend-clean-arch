package taskhandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (t *TaskHandler) Update(c echo.Context) error {

	return c.JSON(http.StatusCreated, echo.Map{
		"status":  true,
		"message": "UPDATE",
		"data":    "",
	})
}
