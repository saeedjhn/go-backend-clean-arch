package taskusecase

type Gateway interface {
	List()
}

type TaskInteractor struct {
	taskGateway Gateway
}

func New(taskGateway Gateway) *TaskInteractor {
	return &TaskInteractor{taskGateway: taskGateway}
}
