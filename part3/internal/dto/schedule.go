package dto

import (
	"part3/internal/model"
	"time"
)

type CreateScheduleRequest struct {
	TaskID  uint      `json:"task_id" binding:"required"`
	StartAt time.Time `json:"start_at" binding:"required"`
	EndAt   time.Time `json:"end_at" binding:"required"`
}

type UpdateScheduleRequest struct {
	StartAt *time.Time `json:"start_at"`
	EndAt   *time.Time `json:"end_at"`
}

type ScheduleResponse struct {
	ID      uint      `json:"id"`
	TaskID  uint      `json:"task_id"`
	StartAt time.Time `json:"start_at"`
	EndAt   time.Time `json:"end_at"`
}

type ListSchedulesResponse struct {
	ID      uint      `json:"id"`
	TaskID  uint      `json:"task_id"`
	StartAt time.Time `json:"start_at"`
	EndAt   time.Time `json:"end_at"`
}

func (r *CreateScheduleRequest) ToModel() *model.Schedule {
	return &model.Schedule{
		TaskID:  r.TaskID,
		StartAt: r.StartAt,
		EndAt:   r.EndAt,
	}
}

func FromScheduleModel(s *model.Schedule) *ScheduleResponse {
	return &ScheduleResponse{
		ID:      s.ID,
		TaskID:  s.TaskID,
		StartAt: s.StartAt,
		EndAt:   s.EndAt,
	}
}

func FromScheduleModelList(schedules []model.Schedule) []ListSchedulesResponse {
	response := make([]ListSchedulesResponse, 0, len(schedules))
	for _, s := range schedules {
		response = append(response, ListSchedulesResponse{
			ID:      s.ID,
			TaskID:  s.TaskID,
			StartAt: s.StartAt,
			EndAt:   s.EndAt,
		})
	}
	return response
}
