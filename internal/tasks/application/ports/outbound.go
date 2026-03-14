package ports

import (
	"context"
	"task_flow_m2_go/internal/tasks/domain"
)

type TaskRepository interface {
	Create(ctx context.Context, task domain.Task) (domain.Task, error)
	GetByProjectID(ctx context.Context, projectID string) (domain.Task, error)
	UpdateStatus(ctx context.Context, projectID string, taskID string, status domain.TaskStatus) (domain.Task, error)
	GetByID(ctx context.Context, projectID, taskID string) (domain.Task, error)
}

type OutboxRepository interface {
	Create(ctx context.Context, event *domain.OutboxEvent) error
}
