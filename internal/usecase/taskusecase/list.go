package taskusecase

import "log"

func (t *TaskInteractor) List() {
	t.repository.List()

	log.Print("TaskInteractor -> List - IMPL ME")
}
