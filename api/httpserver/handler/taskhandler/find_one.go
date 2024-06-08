package taskhandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (t *TaskHandler) FindOne(c echo.Context) error {

	return c.JSON(http.StatusCreated, echo.Map{
		"status":  true,
		"message": "FIND ONE",
		"data":    "",
	})
}
