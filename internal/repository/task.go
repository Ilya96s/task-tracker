package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/task-tracker/internal/model"
)

type TaskRepository struct {
	db *pgxpool.Pool
}

func NewTaskRepository(db *pgxpool.Pool) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (r *TaskRepository) Create(ctx context.Context, name string) (model.Task, error) {
	var task model.Task

	err := r.db.QueryRow(
		ctx,
		`INSERT INTO tasks (name, created_at) VALUES ($1, NOW()) RETURNING id, name, created_at`,
		name,
	).Scan(&task.ID, &task.Name, &task.CreatedAt)

	return task, err
}
