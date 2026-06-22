package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/task-tracker/internal/model"
)

type TaskListRepository struct {
	db *pgxpool.Pool
}

func NewTaskListRepository(db *pgxpool.Pool) *TaskListRepository {
	return &TaskListRepository{
		db: db,
	}
}

func (r *TaskListRepository) Create(ctx context.Context, name string) (model.TaskList, error) {
	var taskList model.TaskList

	err := r.db.QueryRow(
		ctx,
		`INSERT INTO task_lists (name, created_at) VALUES ($1, NOW()) RETURNING id, name, created_at`,
		name,
	).Scan(&taskList.ID, &taskList.Name, &taskList.CreatedAt)

	return taskList, err
}

func (r *TaskListRepository) GetByID(ctx context.Context, id int) (model.TaskList, error) {
	var taskList model.TaskList

	err := r.db.QueryRow(
		ctx,
		`SELECT id, name, created_at, updated_at, tasks FROM task_lists WHERE id = $1`,
		id,
	).Scan(&taskList.ID, &taskList.Name, &taskList.CreatedAt, &taskList.UpdatedAt, &taskList.Tasks)

	return taskList, err
}
