package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

var (
    green = color.New(color.FgGreen).SprintFunc()
    red   = color.New(color.FgRed).SprintFunc()
)

func ListTodos(todos []Todo) {
    if len(todos) == 0 {
        fmt.Println("No todos yet!")
        return
    }
    for i, t := range todos {
        status := "[ ]"
        if t.Done {
            status = green("[x]")
        }
        line := fmt.Sprintf("%d. %s %s", i+1, status, t.Task)
        if t.Deadline != "" {
            line += fmt.Sprintf(" - due %s", t.Deadline)
        }
        fmt.Println(line)
    }
}

func containsIgnoreCase(s, substr string) bool {
    return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}


// Todo represents a single task
type Todo struct {
    Task     string `json:"task"`
    Done     bool   `json:"done"`
    Category string `json:"category"`
    Deadline string `json:"deadline"` // format: YYYY-MM-DD
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
// func ListTodos(todos []Todo) {
//     if len(todos) == 0 {
//         fmt.Println("No todos yet!")
//         return
//     }
//     for i, t := range todos {
//         status := "[ ]"
//         if t.Done {
//             status = "[x]"
//         }
//         cat := ""
//         if t.Category != "" {
//             cat = fmt.Sprintf(" (%s)", t.Category)
//         }
//         dl := ""
//         if t.Deadline != "" {
//             dl = fmt.Sprintf(" - due %s", t.Deadline)
//         }
//         fmt.Printf("%d. %s %s%s%s\n", i+1, status, t.Task, cat, dl)
//     }
// }

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
		category := ""
		deadline := ""
	
		if len(os.Args) >= 4 {
			category = os.Args[3]
		}
		if len(os.Args) >= 5 {
			deadline = os.Args[4]
		}
	
		// Remove existing duplicates of this exact task (task + category + deadline)
		for _, t := range todos {
			if t.Task == task && t.Category == category && t.Deadline == deadline {
				fmt.Println("Task already exists:", task)
				return
			}
		}
	
		// Append new todo and save
		todos = append(todos, Todo{
			Task:     task,
			Category: category,
			Deadline: deadline,
		})
	
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

	case "edit":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo edit <task number> [new task] [category] [deadline]")
			return
		}
	
		index, err := strconv.Atoi(os.Args[2])
		if err != nil || index < 1 || index > len(todos) {
			fmt.Println("Invalid task number.")
			return
		}
	
		task := todos[index-1]
	
		if len(os.Args) >= 4 {
			task.Task = os.Args[3]
		}
		if len(os.Args) >= 5 {
			task.Category = os.Args[4]
		}
		if len(os.Args) >= 6 {
			task.Deadline = os.Args[5]
		}
	
		todos[index-1] = task
		if err := SaveTodos(todos); err != nil {
			fmt.Println("Error saving todos:", err)
			return
		}
	
		fmt.Println("Edited task:", task.Task)

	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo search <keyword>")
			return
		}
	
		keyword := os.Args[2]
		found := false
	
		for i, t := range todos {
			if containsIgnoreCase(t.Task, keyword) || containsIgnoreCase(t.Category, keyword) || containsIgnoreCase(t.Deadline, keyword) {
				fmt.Printf("%d. [%v] %s (%s) - due %s\n", i+1, t.Done, t.Task, t.Category, t.Deadline)
				found = true
			}
		}
	
		if !found {
			fmt.Println("No matching tasks found for:", keyword)
		}
	
	case "filter":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo filter <category|done|pending>")
			return
		}
	
		filter := strings.ToLower(os.Args[2])
		found := false
	
		for i, t := range todos {
			if filter == "done" && t.Done {
				fmt.Printf("%d. [x] %s (%s) - due %s\n", i+1, t.Task, t.Category, t.Deadline)
				found = true
			} else if filter == "pending" && !t.Done {
				fmt.Printf("%d. [ ] %s (%s) - due %s\n", i+1, t.Task, t.Category, t.Deadline)
				found = true
			} else if strings.ToLower(t.Category) == filter {
				fmt.Printf("%d. [%v] %s (%s) - due %s\n", i+1, t.Done, t.Task, t.Category, t.Deadline)
				found = true
			}
		}
	
		if !found {
			fmt.Println("No tasks found for filter:", filter)
		}
	

	default:
		fmt.Println("Unknown command:", command)
	}
}
