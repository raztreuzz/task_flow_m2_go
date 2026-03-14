package dto

import "time"

type CreateTaskRequest struct {
	ProjectID   uint64 `json:"project_id"`
	Category    string `json:"category_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTaskStatusRequest struct {
	Status string `json:"status"`
}

type TaskResponse struct {
	ID          uint64    `json:"id"`
	ProjectID   uint64    `json:"project_id"`
	CategoryID  uint64    `json:"category_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
