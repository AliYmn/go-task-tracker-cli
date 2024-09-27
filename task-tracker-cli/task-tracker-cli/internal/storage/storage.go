package storage

import "task-tracker-cli/internal/task"

type Storage interface {
    SaveTasks(tasks []task.Task) error
    LoadTasks() ([]task.Task, error)
}
