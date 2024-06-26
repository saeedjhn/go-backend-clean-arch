package mocks

import (
	"github.com/golang/mock/gomock"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/servicedto/usertaskservicedto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
	"github.com/saeedjhn/go-backend-clean-arch/internal/service/taskservice"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := NewMockRepository(ctrl)

	service := taskservice.New(mockRepository)

	createTask := entity.Task{
		UserID:      1,
		Title:       "My title",
		Description: "My description",
		Status:      entity.Pending,
	}

	createdTask := entity.Task{
		ID:          1,
		UserID:      1,
		Title:       "My title",
		Description: "My description",
		Status:      entity.Pending,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	mockRepository.EXPECT().Create(createTask).Return(createdTask, nil)
	mockRepository.EXPECT().IsExistsUser(createTask.UserID).Return(true, nil)

	task, err := service.Create(usertaskservicedto.CreateTaskRequest{
		UserID:      createTask.UserID,
		Title:       createTask.Title,
		Description: createTask.Description,
		Status:      createTask.Status,
	})

	assert.NoError(t, err)
	assert.NotNil(t, task.Task)
	assert.Equal(t, uint(1), task.Task.ID)
	assert.Equal(t, "My title", createdTask.Title)

	// More
}
