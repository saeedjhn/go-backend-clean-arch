package mockstest_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/servicedto/usertaskservicedto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
	"github.com/saeedjhn/go-backend-clean-arch/internal/service/taskservice"
	mockstest "github.com/saeedjhn/go-backend-clean-arch/internal/service/taskservice/mocks_test"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mockstest.NewMockRepository(ctrl)

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

	// The difference between assert and require in Testify

	// assert: Checks the condition, but if it fails, test execution continues.
	// That is, if the condition is not checked, the test will continue.
	// require: Checks the condition and stops the test immediately if it fails.

	// Why is it recommended to use require for errors?
	// In many cases, when an error occurs, the point of continuing to run the test is lost
	// because the main condition of the test is not properly established.
	// Using require in these cases will stop the test in the event of an error and
	// prevent the execution of the rest of the test codes.
	// This will help you avoid getting redundant and confusing error messages.۰

	require.NoError(t, err)
	require.NotNil(t, task.Task)
	require.Equal(t, uint(1), task.Task.ID)
	require.Equal(t, "My title", createdTask.Title)

	// More
}
