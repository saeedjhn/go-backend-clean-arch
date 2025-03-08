package bootstrap

import (
	"github.com/saeedjhn/go-domain-driven-design/configs"
	mysqltask "github.com/saeedjhn/go-domain-driven-design/internal/repository/mysql/task"
	mysqluser "github.com/saeedjhn/go-domain-driven-design/internal/repository/mysql/user"
	redisuser "github.com/saeedjhn/go-domain-driven-design/internal/repository/redis/user"
	contract2 "github.com/saeedjhn/go-domain-driven-design/internal/sharedkernel/contract"
	authusecase "github.com/saeedjhn/go-domain-driven-design/internal/usecase/authentication"
	taskusecase "github.com/saeedjhn/go-domain-driven-design/internal/usecase/task"
	userusecase "github.com/saeedjhn/go-domain-driven-design/internal/usecase/user"
	uservalidator "github.com/saeedjhn/go-domain-driven-design/internal/validator/user"
)

type Usecase struct {
	AuthIntr *authusecase.Interactor
	UserIntr *userusecase.Interactor
	TaskIntr *taskusecase.Interactor
}

func NewUsecase(
	config *configs.Config,
	_ contract2.Logger,
	trc contract2.Tracer,
	cache Cache,
	db DB,
) *Usecase {
	var (
		taskRepo    = mysqltask.New(db.MySQL)
		userRepo    = mysqluser.New(trc, db.MySQL)
		userRdsRepo = redisuser.New(cache.Redis) // Or userInMemRepo := inmemoryuser.New(cache.InMem)
	)

	var (
		taskIntr = taskusecase.New(config, taskRepo)
		authIntr = authusecase.New(config.Auth)
		userVld  = uservalidator.New(config.Application.EntropyPassword)
		userIntr = userusecase.New(config, trc, authIntr, userVld, userRdsRepo, userRepo)
	)

	return &Usecase{
		AuthIntr: authIntr,
		UserIntr: userIntr,
		TaskIntr: taskIntr,
	}
}
