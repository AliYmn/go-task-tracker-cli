package task

import "time"

type Task struct {
	ID        int
	Title     string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
