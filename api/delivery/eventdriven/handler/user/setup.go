package user

import (
	"github.com/saeedjhn/go-backend-clean-arch/api/delivery/eventdriven/registery"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	outboxevent "github.com/saeedjhn/go-backend-clean-arch/internal/repository/mysql/outbox_event"
	usermysql "github.com/saeedjhn/go-backend-clean-arch/internal/repository/mysql/user"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/events"
	authusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/authentication"
	outboxusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/outbox"
	userusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/user"
	uservalidator "github.com/saeedjhn/go-backend-clean-arch/internal/validator/user"
)

func Setup(app *bootstrap.Application, er *registery.R) {
	// Dependencies
	repo := usermysql.New(app.Trc, app.MySQL)
	outboxRepo := outboxevent.New(app.MySQL)
	// rdsRepo := userredis.New(app.Cache.Redis) // Or userInMemRepo := inmemoryuser.New(cache.InMem)

	vld := uservalidator.New(app.Config.Application.EntropyPassword)

	authIntr := authusecase.New(app.Config.Auth)
	outboxIntr := outboxusecase.New(app.Config, app.Trc, outboxRepo)
	userIntr := userusecase.New(app.Config, app.Trc, authIntr, outboxIntr, vld, repo)

	handler := New(app.Trc, userIntr)

	er.Register(events.UsersRegistered, handler.Registered)
}
