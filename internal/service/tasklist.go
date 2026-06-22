package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/task-tracker/internal/model"
	"github.com/task-tracker/internal/repository"
)

type TaskListService struct {
	repo *repository.TaskListRepository
}

func NewTaskListService(repo *repository.TaskListRepository) *TaskListService {
	return &TaskListService{
		repo: repo,
	}
}

func (s *TaskListService) Create(ctx context.Context, name string) (model.TaskList, error) {
	return s.repo.Create(ctx, name)
}

func (s *TaskListService) GetByID(ctx context.Context, id int) (model.TaskList, error) {
	tasklist, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return tasklist, fmt.Errorf("task list not found with id %d: %w", id, err)
		}
		return tasklist, fmt.Errorf("unexpected error: %w", err)
	}

	return tasklist, nil
}
