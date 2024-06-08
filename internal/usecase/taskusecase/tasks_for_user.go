package taskusecase

import (
	"go-backend-clean-arch/internal/domain"
	"time"
)

func (t *TaskInteractor) TasksForUser(userID uint) ([]domain.Task, error) {
	// Fetch from DB
	return []domain.Task{{
		ID:          1,
		UserID:      1,
		Title:       "T1",
		Description: "D1",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, {
		ID:          2,
		UserID:      2,
		Title:       "T2",
		Description: "D2",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}}, nil
}
