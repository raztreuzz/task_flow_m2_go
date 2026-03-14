package ports

import (
	"context"
	"task_flow_m2_go/internal/tasks/application/dto"
	"task_flow_m2_go/internal/tasks/domain"
)

type TaskUseCase interface {
	CreateTask(ctx context.Context, request dto.CreateTaskRequest) (*domain.Task, error)
	GetTaskByProjectID(ctx context.Context, projectID uint64) ([]dto.TaskResponse, error)
	updateTaskStatus(ctx context.Context, taskID uint64, request dto.UpdateTaskStatusRequest) (*dto.TaskResponse, error)
}
