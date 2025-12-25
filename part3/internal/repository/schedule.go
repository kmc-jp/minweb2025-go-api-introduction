package repository

import (
	"part3/internal/model"

	"gorm.io/gorm"
)

type ScheduleRepository interface {
	Create(schedule *model.Schedule) error
	FindByID(id uint) (*model.Schedule, error)
	FindByTaskID(taskID uint) ([]model.Schedule, error)
	Update(schedule *model.Schedule) error
	Delete(schedule *model.Schedule) error
	List() ([]model.Schedule, error)
}

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &scheduleRepository{db: db}
}

func (r *scheduleRepository) Create(schedule *model.Schedule) error {
	return r.db.Create(schedule).Error
}

func (r *scheduleRepository) FindByID(id uint) (*model.Schedule, error) {
	var schedule model.Schedule
	if err := r.db.First(&schedule, id).Error; err != nil {
		return nil, err
	}
	return &schedule, nil
}

func (r *scheduleRepository) FindByTaskID(taskID uint) ([]model.Schedule, error) {
	var schedules []model.Schedule
	if err := r.db.Where("task_id = ?", taskID).Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}

func (r *scheduleRepository) Update(schedule *model.Schedule) error {
	return r.db.Save(schedule).Error
}

func (r *scheduleRepository) Delete(schedule *model.Schedule) error {
	return r.db.Delete(schedule).Error
}

func (r *scheduleRepository) List() ([]model.Schedule, error) {
	var schedules []model.Schedule
	if err := r.db.Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}
