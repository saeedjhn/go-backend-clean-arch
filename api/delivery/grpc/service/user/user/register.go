package user

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	outboxevent "github.com/saeedjhn/go-backend-clean-arch/internal/repository/mysql/outbox_event"
	mysqluser "github.com/saeedjhn/go-backend-clean-arch/internal/repository/mysql/user"
	pb "github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/proto/user/gen"
	authusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/authentication"
	outboxusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/outbox"
	userusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/user"
	uservalidator "github.com/saeedjhn/go-backend-clean-arch/internal/validator/user"
	"google.golang.org/grpc"
)

func Register(app *bootstrap.Application, gs grpc.ServiceRegistrar) {
	// Way 1
	// us := New(app.Usecase.UserIntr)

	// Way 2
	// Dependencies
	repo := mysqluser.New(app.Trc, app.MySQL)
	outboxRepo := outboxevent.New(app.MySQL)
	// rdsRepo := redisuser.New(app.Cache.Redis) // Or userInMemRepo := inmemoryuser.New(cache.InMem)

	vld := uservalidator.New(app.Config.Application.EntropyPassword)

	authIntr := authusecase.New(app.Config.Auth)
	outboxIntr := outboxusecase.New(app.Config, app.Trc, outboxRepo)
	userIntr := userusecase.New(app.Config, app.Trc, authIntr, outboxIntr, vld, repo)

	us := New(userIntr)

	pb.RegisterUserServiceServer(gs, us)
}
