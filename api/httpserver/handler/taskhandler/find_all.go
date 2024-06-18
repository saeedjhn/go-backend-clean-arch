package taskhandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (t *TaskHandler) FindAll(c echo.Context) error {

	return c.JSON(http.StatusOK, echo.Map{
		"status":  true,
		"message": "FIND ALL",
		"data":    "",
	})
}
