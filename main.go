package main

import (
	"fmt"     // Package for formatted I/O
	"os"      // Package for interacting with the operating system
	"strconv" // Package for converting strings to other types
	"time"    // Package for working with time
)

// main is the entry point of the application
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <command> [arguments]") // Print usage if no command is provided
		return
	}

	command := os.Args[1] // Get the command from the arguments

	// Create a channel to handle task operations
	taskChan := make(chan string)

	// Use a Go routine to handle the command
	go func() {
		switch command {
		case "add":
			if len(os.Args) < 3 {
				taskChan <- "Usage: task-cli add <description>" // Print usage if description is missing
				return
			}
			description := os.Args[2]
			addTask(description, taskChan)
		case "list":
			listTasks(taskChan)
		case "update":
			if len(os.Args) < 4 {
				taskChan <- "Usage: task-cli update <id> <description>" // Print usage if id or description is missing
				return
			}
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				taskChan <- "Invalid ID"
				return
			}
			description := os.Args[3]
			updateTask(id, description, taskChan)
		case "delete":
			if len(os.Args) < 3 {
				taskChan <- "Usage: task-cli delete <id>" // Print usage if id is missing
				return
			}
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				taskChan <- "Invalid ID"
				return
			}
			deleteTask(id, taskChan)
		case "mark-in-progress":
			if len(os.Args) < 3 {
				taskChan <- "Usage: task-cli mark-in-progress <id>" // Print usage if id is missing
				return
			}
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				taskChan <- "Invalid ID"
				return
			}
			markTask(id, "in-progress", taskChan)
		case "mark-done":
			if len(os.Args) < 3 {
				taskChan <- "Usage: task-cli mark-done <id>" // Print usage if id is missing
				return
			}
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				taskChan <- "Invalid ID"
				return
			}
			markTask(id, "done", taskChan)
		default:
			taskChan <- fmt.Sprintf("Unknown command: %s", command) // Print error if command is unknown
		}
	}()

	// Print the result from the channel
	fmt.Println(<-taskChan)
}

// addTask adds a new task with the given description
func addTask(description string, taskChan chan string) {
	tasks, err := loadTasks() // Load existing tasks
	if err != nil {
		taskChan <- fmt.Sprintf("Error loading tasks: %v", err)
		return
	}

	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1 // Generate a new ID
	}

	task := &Task{
		ID:          id,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, task) // Add the new task to the list

	if err := saveTasks(tasks); err != nil {
		taskChan <- fmt.Sprintf("Error saving task: %v", err)
		return
	}

	taskChan <- fmt.Sprintf("Task added successfully (ID: %d)", id)
}

// listTasks lists all tasks
func listTasks(taskChan chan string) {
	tasks, err := loadTasks() // Load existing tasks
	if err != nil {
		taskChan <- fmt.Sprintf("Error loading tasks: %v", err)
		return
	}

	result := ""
	for _, task := range tasks {
		result += fmt.Sprintf("ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n",
			task.ID, task.Description, task.Status, task.CreatedAt.Format(time.RFC3339), task.UpdatedAt.Format(time.RFC3339))
	}

	taskChan <- result
}

// updateTask updates the description of a task with the given ID
func updateTask(id int, description string, taskChan chan string) {
	tasks, err := loadTasks() // Load existing tasks
	if err != nil {
		taskChan <- fmt.Sprintf("Error loading tasks: %v", err)
		return
	}

	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		taskChan <- "Task not found"
		return
	}

	if err := saveTasks(tasks); err != nil {
		taskChan <- fmt.Sprintf("Error saving task: %v", err)
		return
	}

	taskChan <- "Task updated successfully"
}

// deleteTask deletes a task with the given ID
func deleteTask(id int, taskChan chan string) {
	tasks, err := loadTasks() // Load existing tasks
	if err != nil {
		taskChan <- fmt.Sprintf("Error loading tasks: %v", err)
		return
	}

	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		taskChan <- "Task not found"
		return
	}

	if err := saveTasks(tasks); err != nil {
		taskChan <- fmt.Sprintf("Error saving task: %v", err)
		return
	}

	taskChan <- "Task deleted successfully"
}

// markTask updates the status of a task with the given ID
func markTask(id int, status string, taskChan chan string) {
	tasks, err := loadTasks() // Load existing tasks
	if err != nil {
		taskChan <- fmt.Sprintf("Error loading tasks: %v", err)
		return
	}

	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		taskChan <- "Task not found"
		return
	}

	if err := saveTasks(tasks); err != nil {
		taskChan <- fmt.Sprintf("Error saving task: %v", err)
		return
	}

	taskChan <- fmt.Sprintf("Task marked as %s successfully", status)
}
