package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Todo represents a single task
type Todo struct {
	Task string
	Done bool
}

// File to store todos
const todoFile = "todos.json"

// LoadTodos reads todos from JSON file
func LoadTodos() ([]Todo, error) {
	file, err := os.ReadFile(todoFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Todo{}, nil
		}
		return nil, err
	}

	var todos []Todo
	if err := json.Unmarshal(file, &todos); err != nil {
		return nil, err
	}
	return todos, nil
}

// SaveTodos writes todos to JSON file
func SaveTodos(todos []Todo) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(todoFile, data, 0644)
}

// Helper to print todos
func ListTodos(todos []Todo) {
	if len(todos) == 0 {
		fmt.Println("No todos yet!")
		return
	}
	for i, t := range todos {
		status := "[ ]"
		if t.Done {
			status = "[x]"
		}
		fmt.Printf("%d. %s %s\n", i+1, status, t.Task)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [add|list|done|delete] [task]")
		return
	}

	command := os.Args[1]
	todos, err := LoadTodos()
	if err != nil {
		fmt.Println("Error loading todos:", err)
		return
	}

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task.")
			return
		}
		task := os.Args[2]
		todos = append(todos, Todo{Task: task})
		if err := SaveTodos(todos); err != nil {
			fmt.Println("Error saving todos:", err)
			return
		}
		fmt.Println("Added:", task)

	case "list":
		ListTodos(todos)

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Please provide the todo number.")
			return
		}
		index, _ := strconv.Atoi(os.Args[2])
		if index < 1 || index > len(todos) {
			fmt.Println("Invalid todo number.")
			return
		}
		todos[index-1].Done = true
		if err := SaveTodos(todos); err != nil {
			fmt.Println("Error saving todos:", err)
			return
		}
		fmt.Println("Marked as done:", todos[index-1].Task)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide the todo number.")
			return
		}
		index, _ := strconv.Atoi(os.Args[2])
		if index < 1 || index > len(todos) {
			fmt.Println("Invalid todo number.")
			return
		}
		deleted := todos[index-1].Task
		todos = append(todos[:index-1], todos[index:]...)
		if err := SaveTodos(todos); err != nil {
			fmt.Println("Error saving todos:", err)
			return
		}
		fmt.Println("Deleted:", deleted)

	default:
		fmt.Println("Unknown command:", command)
	}
}
