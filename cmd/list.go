package cmd

import (
	"project-app-todo-list-cli-rafli-nur-rahman/service"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := service.GetAllTasks()

		if len(tasks) == 0 {
			println("No tasks found.")
			return
		}

		displayTasks(tasks)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
