package main

import (
	"task-tracker-cli/internal/storage"
	"task-tracker-cli/internal/task"
	"task-tracker-cli/pkg/cli"
)

func main() {
    // Initialize TaskManager
    taskManager := &task.TaskManager{}

    // Set up storage
    storage := &storage.JSONStorage{FilePath: "tasks.json"}

    // Initialize CLI with dependencies
    cli := &cli.CLI{
        TaskManager: taskManager,
        Storage:     storage,
    }

    // Run the CLI
    cli.Run()
}
