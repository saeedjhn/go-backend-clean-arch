package taskinggateway

type TaskInteractor interface {
	List()
}

type TaskingGateway struct {
	taskInteractor TaskInteractor
}

func New(taskInteractor TaskInteractor) *TaskingGateway {
	return &TaskingGateway{taskInteractor: taskInteractor}
}
