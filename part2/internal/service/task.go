package service

import (
	"part2/internal/dto"
	"part2/internal/model"
	"part2/internal/repository"
)

type TaskService interface {
	CreateTask(req *dto.CreateTaskRequest) (*model.Task, error)
	GetTaskByID(id uint) (*model.Task, error)
	UpdateTask(id uint, req *dto.UpdateTaskRequest) (*model.Task, error)
	DeleteTask(id uint) error
	ListTasks() ([]model.Task, error)
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(req *dto.CreateTaskRequest) (*model.Task, error) {
	task := &model.Task{
		Title:       req.Title,
		Description: req.Description,
		Completed:   false,
	}
	if err := s.repo.Create(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (s *taskService) GetTaskByID(id uint) (*model.Task, error) {
	return s.repo.FindByID(id)
}

func (s *taskService) UpdateTask(id uint, req *dto.UpdateTaskRequest) (*model.Task, error) {
	task, err := s.repo.FindByID(id)
	if err != nil {
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
	return task, nil
}

func (s *taskService) DeleteTask(id uint) error {
	task, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(task)
}

func (s *taskService) ListTasks() ([]model.Task, error) {
	return s.repo.List()
}
