package user

import (
	pb "github.com/saeedjhn/go-backend-clean-arch/api/proto/user/gen"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	mysqluser "github.com/saeedjhn/go-backend-clean-arch/internal/repository/mysql/user"
	redisuser "github.com/saeedjhn/go-backend-clean-arch/internal/repository/redis/user"
	authusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/auth"
	userusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/user"
	uservalidator "github.com/saeedjhn/go-backend-clean-arch/internal/validator/user"
	"google.golang.org/grpc"
)

func Register(app *bootstrap.Application, gs grpc.ServiceRegistrar) {
	// Way 1
	// us := New(app.Usecase.UserIntr)

	// Way 2
	// Dependencies
	repo := mysqluser.New(app.Trc, app.DB.MySQL)
	rdsRepo := redisuser.New(app.Cache.Redis) // Or userInMemRepo := inmemoryuser.New(cache.InMem)

	vld := uservalidator.New(app.Config.Application.EntropyPassword)

	authIntr := authusecase.New(app.Config.Auth)
	userIntr := userusecase.New(app.Config, app.Trc, vld, authIntr, rdsRepo, repo)

	us := New(userIntr)

	pb.RegisterUserServiceServer(gs, us)
}
