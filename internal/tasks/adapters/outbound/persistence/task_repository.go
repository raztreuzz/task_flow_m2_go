package persistence

import (
	"context"
	"database/sql"

	"task_flow_m2_go/internal/tasks/domain"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (r *TaskRepository) Create(ctx context.Context, task *domain.Task) error {
	query := `
		INSERT INTO tasks (project_id, category_id, title, description, status, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		task.ProjectID,
		task.CategoryID,
		task.Title,
		task.Description,
		task.Status,
		task.CreatedAt,
		task.UpdatedAt,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	task.ID = uint64(id)
	return nil
}

func (r *TaskRepository) GetByProjectID(ctx context.Context, projectID uint64) ([]domain.Task, error) {
	query := `
		SELECT id, project_id, category_id, title, description, status, created_at, updated_at
		FROM tasks
		WHERE project_id = ?
		ORDER BY id DESC
	`

	rows, err := r.db.QueryContext(ctx, query, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]domain.Task, 0)

	for rows.Next() {
		var task domain.Task

		err := rows.Scan(
			&task.ID,
			&task.ProjectID,
			&task.CategoryID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepository) GetByID(ctx context.Context, taskID uint64) (*domain.Task, error) {
	query := `
		SELECT id, project_id, category_id, title, description, status, created_at, updated_at
		FROM tasks
		WHERE id = ?
		LIMIT 1
	`

	var task domain.Task

	err := r.db.QueryRowContext(ctx, query, taskID).Scan(
		&task.ID,
		&task.ProjectID,
		&task.CategoryID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) UpdateStatus(ctx context.Context, taskID uint64, status domain.TaskStatus) error {
	query := `
		UPDATE tasks
		SET status = ?, updated_at = NOW()
		WHERE id = ?
	`

	_, err := r.db.ExecContext(ctx, query, status, taskID)
	return err
}
