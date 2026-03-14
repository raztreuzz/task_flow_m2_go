package services

import (
	"context"
	"encoding/json"
	"time"

	"task_flow_m2_go/internal/tasks/application/dto"
	"task_flow_m2_go/internal/tasks/application/ports"
	"task_flow_m2_go/internal/tasks/domain"
)

type TaskService struct {
	taskRepository   ports.TaskRepository
	outboxRepository ports.OutboxRepository
}

func NewTaskService(
	taskRepository ports.TaskRepository,
	outboxRepository ports.OutboxRepository,
) *TaskService {
	return &TaskService{
		taskRepository:   taskRepository,
		outboxRepository: outboxRepository,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, request dto.CreateTaskRequest) (*dto.TaskResponse, error) {
	now := time.Now().UTC()

	task := &domain.Task{
		ProjectID:   request.ProjectID,
		CategoryID:  request.CategoryID,
		Title:       request.Title,
		Description: request.Description,
		Status:      domain.TaskStatusPending,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	err := s.taskRepository.Create(ctx, task)
	if err != nil {
		return nil, err
	}

	payloadBytes, err := json.Marshal(map[string]any{
		"task_id":     task.ID,
		"project_id":  task.ProjectID,
		"category_id": task.CategoryID,
		"title":       task.Title,
		"description": task.Description,
		"status":      task.Status,
	})
	if err != nil {
		return nil, err
	}

	event := &domain.OutboxEvent{
		Aggregate:   "task",
		AggregateID: task.ID,
		Type:        "task.created",
		Payload:     string(payloadBytes),
		CreatedAt:   now,
		Processed:   false,
	}

	err = s.outboxRepository.Create(ctx, event)
	if err != nil {
		return nil, err
	}

	response := &dto.TaskResponse{
		ID:          task.ID,
		ProjectID:   task.ProjectID,
		CategoryID:  task.CategoryID,
		Title:       task.Title,
		Description: task.Description,
		Status:      string(task.Status),
		CreatedAt:   task.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   task.UpdatedAt.Format(time.RFC3339),
	}

	return response, nil
}

func (s *TaskService) GetTasksByProjectID(ctx context.Context, projectID uint64) ([]dto.TaskResponse, error) {
	tasks, err := s.taskRepository.GetByProjectID(ctx, projectID)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.TaskResponse, 0, len(tasks))
	for _, task := range tasks {
		responses = append(responses, dto.TaskResponse{
			ID:          task.ID,
			ProjectID:   task.ProjectID,
			CategoryID:  task.CategoryID,
			Title:       task.Title,
			Description: task.Description,
			Status:      string(task.Status),
			CreatedAt:   task.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   task.UpdatedAt.Format(time.RFC3339),
		})
	}

	return responses, nil
}

func (s *TaskService) UpdateTaskStatus(ctx context.Context, taskID uint64, request dto.UpdateTaskStatusRequest) (*dto.TaskResponse, error) {
	task, err := s.taskRepository.GetByID(ctx, taskID)
	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, domain.ErrTaskNotFound
	}

	newStatus := domain.TaskStatus(request.Status)
	task.Status = newStatus
	task.UpdatedAt = time.Now().UTC()

	err = s.taskRepository.UpdateStatus(ctx, taskID, newStatus)
	if err != nil {
		return nil, err
	}

	payloadBytes, err := json.Marshal(map[string]any{
		"task_id": task.ID,
		"status":  task.Status,
	})
	if err != nil {
		return nil, err
	}

	event := &domain.OutboxEvent{
		Aggregate:   "task",
		AggregateID: task.ID,
		Type:        "task.status_changed",
		Payload:     string(payloadBytes),
		CreatedAt:   task.UpdatedAt,
		Processed:   false,
	}

	err = s.outboxRepository.Create(ctx, event)
	if err != nil {
		return nil, err
	}

	response := &dto.TaskResponse{
		ID:          task.ID,
		ProjectID:   task.ProjectID,
		CategoryID:  task.CategoryID,
		Title:       task.Title,
		Description: task.Description,
		Status:      string(task.Status),
		CreatedAt:   task.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   task.UpdatedAt.Format(time.RFC3339),
	}

	return response, nil
}
