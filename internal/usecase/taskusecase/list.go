package taskusecase

import "log"

func (t *TaskInteractor) List() {
	t.taskGateway.List()

	log.Print("TaskInteractor -> List - IMPL ME")
}
