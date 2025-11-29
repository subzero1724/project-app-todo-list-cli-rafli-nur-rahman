package service

import (
	"errors"
	"fmt"
	"os"
	"project-app-todo-list-cli-rafli-nur-rahman/model"
	"project-app-todo-list-cli-rafli-nur-rahman/utils"
	"strings"
	"text/tabwriter"
	"time"
)

var dataFile = "data/tasks.json"

// Load tasks from JSON file
func load() ([]model.Task, error) {
	var tasks []model.Task

	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		return tasks, nil
	}

	err := utils.LoadJSON(dataFile, &tasks)
	return tasks, err
}

// Save tasks to JSON file
func save(tasks []model.Task) error {
	return utils.SaveJSON(dataFile, tasks)
}

// Add a new task
func AddTask(title, description string) error {
	tasks, _ := load()

	// Avoid duplicate titles
	for _, t := range tasks {
		if strings.EqualFold(t.Title, title) {
			return errors.New("task with the same title already exists")
		}
	}

	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}

	newTask := model.Task{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      "todo",
		Priority:    "normal",
		CreatedAt:   time.Now().Format(time.RFC3339),
	}

	tasks = append(tasks, newTask)
	return save(tasks)
}

// Update task
func UpdateTask(id int, title, desc, status, priority string) error {
	tasks, err := load()
	if err != nil {
		return err
	}

	var found bool

	for i, t := range tasks {
		if t.ID == id {
			found = true

			if title != "" {
				tasks[i].Title = title
			}
			if desc != "" {
				tasks[i].Description = desc
			}
			if status != "" {
				if status != "todo" && status != "done" {
					return errors.New("status must be: todo / done")
				}
				tasks[i].Status = status
			}
			if priority != "" {
				if priority != "low" && priority != "normal" && priority != "high" {
					return errors.New("priority must be: low / normal / high")
				}
				tasks[i].Priority = priority
			}
		}
	}

	if !found {
		return errors.New("task ID not found")
	}

	return save(tasks)
}

// List tasks (all or only todo)
func ListTasks(all bool) {
	tasks, _ := load()

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprint(w, "ID\tTitle\tStatus\tPriority\tCreatedAt\n")

	for _, t := range tasks {
		if !all && t.Status == "done" {
			continue
		}
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n",
			t.ID, t.Title, t.Status, t.Priority, t.CreatedAt)
	}
	w.Flush()
}

// Mark a task as done
func MarkTaskDone(id int) error {
	tasks, err := load()
	if err != nil {
		return err
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Status = "done"
			return save(tasks)
		}
	}

	return errors.New("task not found")
}

// Delete a task
func DeleteTask(id int) error {
	tasks, _ := load()

	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return save(tasks)
		}
	}

	return errors.New("task not found")
}

// NEW — Search tasks by keyword (title or description)
func SearchTasks(keyword string) []model.Task {
	tasks, _ := load()

	var results []model.Task
	key := strings.ToLower(keyword)

	for _, t := range tasks {
		if strings.Contains(strings.ToLower(t.Title), key) ||
			strings.Contains(strings.ToLower(t.Description), key) {
			results = append(results, t)
		}
	}

	return results
}

// Return all tasks
func GetAllTasks() []model.Task {
	tasks, err := load()
	if err != nil {
		return []model.Task{}
	}
	return tasks
}
