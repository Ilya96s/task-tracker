package model

import "time"

type Task struct {
	ID         int
	TaskListID int
	Name       string
	Done       bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
