package user

import (
	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	outboxevent "github.com/saeedjhn/go-backend-clean-arch/internal/repository/mysql/outbox_event"
	usermysql "github.com/saeedjhn/go-backend-clean-arch/internal/repository/mysql/user"
	authusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/authentication"
	outboxusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/outbox"
	userusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/user"
	uservalidator "github.com/saeedjhn/go-backend-clean-arch/internal/validator/user"
)

func Setup(app *bootstrap.Application, e *echo.Echo) {
	// Dependencies
	repo := usermysql.New(app.Trc, app.MySQL)
	outboxRepo := outboxevent.New(app.MySQL)
	// rdsRepo := userredis.New(app.Cache.Redis) // Or userInMemRepo := inmemoryuser.New(cache.InMem)

	vld := uservalidator.New(app.Config.Application.EntropyPassword)

	authIntr := authusecase.New(app.Config.Auth)
	outboxIntr := outboxusecase.New(app.Config, app.Trc, outboxRepo)
	userIntr := userusecase.New(app.Config, app.Trc, authIntr, outboxIntr, vld, repo)

	handler := New(app.Trc, authIntr, userIntr)

	// Way 1
	handler.SetRoutes(e)

	// Way 2
	// group := e.Group("/users")
	// {
	// 	publicG := group.Group("")
	// 	{
	// 		publicG.POST("/refresh-token", handler.RefreshToken)
	// 	}
	//
	// 	authG := group.Group("/auth")
	// 	{
	// 		authG.POST("/register", handler.Register)
	// 		authG.POST("/login", handler.Login)
	// 	}
	//
	// 	protectedG := group.Group("")
	// 	protectedG.Use(mymiddleware.Auth(authIntr))
	// 	{
	// 		protectedG.GET("/profile", handler.Profile)
	// 	}
	// }
}
