package cmd

import (
	"os"
	"project-app-todo-list-cli-rafli-nur-rahman/service"
	"project-app-todo-list-cli-rafli-nur-rahman/utils"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := service.GetAllTasks()

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)

		t.AppendHeader(table.Row{"No", "Task", "Status", "Priority"})

		for i, task := range tasks {
			coloredStatus := utils.ColorStatus(task.Status)
			t.AppendRow(table.Row{i + 1, task.Title, coloredStatus, task.Priority})
		}

		t.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
