package task

import (
	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	mysqltask "github.com/saeedjhn/go-backend-clean-arch/internal/repository/mysql/task"
	authusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/auth"
	taskusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/task"
)

func Setup(app *bootstrap.Application, e *echo.Echo) {
	// Dependencies
	repo := mysqltask.New(app.DB.MySQL)

	authIntr := authusecase.New(app.Config.Auth)
	taskIntr := taskusecase.New(app.Config, repo)

	New(app.Trc, authIntr, taskIntr).SetRoutes(e)
}
