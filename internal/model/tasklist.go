package model

import "time"

type TaskList struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Tasks     []Task
}
