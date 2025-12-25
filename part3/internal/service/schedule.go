package service

import (
	"errors"

	"part3/internal/dto"
	"part3/internal/repository"

	"gorm.io/gorm"
)

var (
	ErrScheduleNotFound = errors.New("schedule not found")
)

type ScheduleService interface {
	CreateSchedule(req *dto.CreateScheduleRequest) (*dto.ScheduleResponse, error)
	GetScheduleByID(id uint) (*dto.ScheduleResponse, error)
	GetSchedulesByTaskID(taskID uint) ([]dto.ListSchedulesResponse, error)
	UpdateSchedule(id uint, req *dto.UpdateScheduleRequest) (*dto.ScheduleResponse, error)
	DeleteSchedule(id uint) error
	ListSchedules() ([]dto.ListSchedulesResponse, error)
}

type scheduleService struct {
	repo     repository.ScheduleRepository
	taskRepo repository.TaskRepository
}

func NewScheduleService(repo repository.ScheduleRepository, taskRepo repository.TaskRepository) ScheduleService {
	return &scheduleService{
		repo:     repo,
		taskRepo: taskRepo,
	}
}

func (s *scheduleService) CreateSchedule(req *dto.CreateScheduleRequest) (*dto.ScheduleResponse, error) {
	// タスクの存在確認
	_, err := s.taskRepo.FindByID(req.TaskID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTaskNotFound
		}
		return nil, err
	}

	schedule := req.ToModel()

	if err := s.repo.Create(schedule); err != nil {
		return nil, err
	}

	return dto.FromScheduleModel(schedule), nil
}

func (s *scheduleService) GetScheduleByID(id uint) (*dto.ScheduleResponse, error) {
	schedule, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrScheduleNotFound
		}
		return nil, err
	}
	return dto.FromScheduleModel(schedule), nil
}

func (s *scheduleService) GetSchedulesByTaskID(taskID uint) ([]dto.ListSchedulesResponse, error) {
	schedules, err := s.repo.FindByTaskID(taskID)
	if err != nil {
		return nil, err
	}
	return dto.FromScheduleModelList(schedules), nil
}

func (s *scheduleService) UpdateSchedule(id uint, req *dto.UpdateScheduleRequest) (*dto.ScheduleResponse, error) {
	schedule, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrScheduleNotFound
		}
		return nil, err
	}

	if req.StartAt != nil {
		schedule.StartAt = *req.StartAt
	}
	if req.EndAt != nil {
		schedule.EndAt = *req.EndAt
	}

	if err := s.repo.Update(schedule); err != nil {
		return nil, err
	}

	return dto.FromScheduleModel(schedule), nil
}

func (s *scheduleService) DeleteSchedule(id uint) error {
	schedule, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrScheduleNotFound
		}
		return err
	}
	return s.repo.Delete(schedule)
}

func (s *scheduleService) ListSchedules() ([]dto.ListSchedulesResponse, error) {
	schedules, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	return dto.FromScheduleModelList(schedules), nil
}
