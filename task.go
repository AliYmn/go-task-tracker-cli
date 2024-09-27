package main

import (
	"encoding/json" // Package for encoding and decoding JSON
	"os"            // Package for interacting with the operating system
	"time"          // Package for working with time
)

// Task struct represents a task with various properties
type Task struct {
	ID          int       `json:"id"`          // Unique identifier for the task
	Description string    `json:"description"` // Short description of the task
	Status      string    `json:"status"`      // Status of the task (todo, in-progress, done)
	CreatedAt   time.Time `json:"createdAt"`   // Timestamp when the task was created
	UpdatedAt   time.Time `json:"updatedAt"`   // Timestamp when the task was last updated
}

const (
	filePath = "tasks.json" // Path to the JSON file where tasks are stored
)

// loadTasks reads the tasks from the JSON file and returns them
func loadTasks() ([]*Task, error) {
	var tasks []*Task
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return tasks, nil // Return an empty list if the file does not exist
	}
	data, err := os.ReadFile(filePath) // Read the file
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tasks) // Decode JSON data into tasks
	return tasks, err
}

// saveTasks writes the tasks to the JSON file
func saveTasks(tasks []*Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ") // Encode tasks to JSON
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644) // Write JSON data to file
}
