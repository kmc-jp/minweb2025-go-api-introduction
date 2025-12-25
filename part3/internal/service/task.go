package service

import (
	"errors"

	"part2/internal/dto"
	"part2/internal/repository"

	"gorm.io/gorm"
)

var (
	ErrTaskNotFound = errors.New("task not found")
)

type TaskService interface {
	CreateTask(req *dto.CreateTaskRequest) (*dto.TaskResponse, error)
	GetTaskByID(id uint) (*dto.TaskResponse, error)
	UpdateTask(id uint, req *dto.UpdateTaskRequest) (*dto.TaskResponse, error)
	DeleteTask(id uint) error
	ListTasks() ([]dto.ListTasksResponse, error)
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(req *dto.CreateTaskRequest) (*dto.TaskResponse, error) {

	task := req.ToModel()

	if err := s.repo.Create(task); err != nil {
		return nil, err
	}

	return dto.FromModel(task), nil
}

func (s *taskService) GetTaskByID(id uint) (*dto.TaskResponse, error) {
	task, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTaskNotFound
		}
		return nil, err
	}
	return dto.FromModel(task), nil
}

func (s *taskService) UpdateTask(id uint, req *dto.UpdateTaskRequest) (*dto.TaskResponse, error) {
	task, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTaskNotFound
		}
		return nil, err
	}

	if req.Title != nil {
		task.Title = *req.Title
	}
	if req.Description != nil {
		task.Description = *req.Description
	}
	if req.Completed != nil {
		task.Completed = *req.Completed
	}

	if err := s.repo.Update(task); err != nil {
		return nil, err
	}

	return dto.FromModel(task), nil
}

func (s *taskService) DeleteTask(id uint) error {
	task, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrTaskNotFound
		}
		return err
	}
	return s.repo.Delete(task)
}

func (s *taskService) ListTasks() ([]dto.ListTasksResponse, error) {
	tasks, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	return dto.FromModelList(tasks), nil
}
