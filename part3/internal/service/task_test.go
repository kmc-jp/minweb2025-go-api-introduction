package service

import (
	"part3/internal/dto"
	"part3/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateTask(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := repository.NewMockTaskRepository(ctrl)

	mockRepo.EXPECT().
		Create(gomock.Any()).
		Return(nil)

	service := NewTaskService(mockRepo)

	req := &dto.CreateTaskRequest{
		Title:       "Test Task",
		Description: "This is a test task",
	}

	res, err := service.CreateTask(req)

	assert.NoError(t, err)
	assert.Equal(t, req.Title, res.Title)
	assert.Equal(t, req.Description, res.Description)
}
