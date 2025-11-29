package cmd

import (
	"os"
	"project-app-todo-list-cli-rafli-nur-rahman/model"

	"github.com/jedib0t/go-pretty/v6/table"
)

func displayTasks(tasks []model.Task) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(table.Row{"No", "Task", "Status", "Priority", "Created At"})

	for i, task := range tasks {
		t.AppendRow(table.Row{
			i + 1,
			task.Title,
			task.Status,
			task.Priority,
			task.CreatedAt,
		})
	}

	t.Render()
}
