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
		if t.Title == title {
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
		Status:      "todo",   // <-- default
		Priority:    "normal", // <-- default
		CreatedAt:   time.Now().Format(time.RFC3339),
	}

	tasks = append(tasks, newTask)
	return save(tasks)
}

// List tasks (pending only OR all)
func ListTasks(all bool) {
	tasks, _ := load()

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprint(w, "ID\tTitle\tStatus\tPriority\tCreatedAt\n")

	for _, t := range tasks {
		if !all && t.Status == "done" {
			continue
		}
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n", t.ID, t.Title, t.Status, t.Priority, t.CreatedAt)
	}
	w.Flush()
}

// Mark task as done
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

// Search tasks by keyword
func SearchTasks(keyword string) {
	tasks, _ := load()
	keyword = strings.ToLower(keyword)

	for _, t := range tasks {
		if strings.Contains(strings.ToLower(t.Title), keyword) {
			fmt.Println("- " + t.Title)
		}
	}
}

// Return all tasks as []model.Task
func GetAllTasks() []model.Task {
	tasks, err := load()
	if err != nil {
		return []model.Task{}
	}
	return tasks
}
