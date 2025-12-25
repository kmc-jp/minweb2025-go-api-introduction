package dto

import (
	"part2/internal/model"
)

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type UpdateTaskRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Completed   *bool   `json:"completed"`
}

type TaskResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type ListTasksResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

func (r *CreateTaskRequest) ToModel() *model.Task {
	return &model.Task{
		Title:       r.Title,
		Description: r.Description,
		Completed:   false, // デフォルト値
	}
}

func FromModel(t *model.Task) *TaskResponse {
	return &TaskResponse{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Completed:   t.Completed,
	}
}

func FromModelList(tasks []model.Task) []ListTasksResponse {
	response := make([]ListTasksResponse, 0, len(tasks)) // 空スライスで初期化
	for _, t := range tasks {
		response = append(response, ListTasksResponse{
			ID:    t.ID,
			Title: t.Title,
		})
	}
	return response
}
