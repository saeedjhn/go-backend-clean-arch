package userusecase

type Gateway interface {
	SaveUserToDB()
	TaskList()
}

type UserInteractor struct {
	userGateway Gateway
}

func New(userGateway Gateway) *UserInteractor {
	return &UserInteractor{userGateway: userGateway}
}
