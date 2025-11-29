package cmd

import (
	"fmt"
	"project-app-todo-list-cli-rafli-nur-rahman/service"

	"github.com/spf13/cobra"
)

var (
	updateID       int
	updateTitle    string
	updateDesc     string
	updateStatus   string
	updatePriority string
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task by ID",
	Run: func(cmd *cobra.Command, args []string) {

		if updateID == 0 {
			fmt.Println("Error: --id is required")
			return
		}

		err := service.UpdateTask(updateID, updateTitle, updateDesc, updateStatus, updatePriority)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Task updated successfully!")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().IntVar(&updateID, "id", 0, "Task ID to update")
	updateCmd.Flags().StringVar(&updateTitle, "title", "", "New title")
	updateCmd.Flags().StringVar(&updateDesc, "desc", "", "New description")
	updateCmd.Flags().StringVar(&updateStatus, "status", "", "todo / done")
	updateCmd.Flags().StringVar(&updatePriority, "priority", "", "low / normal / high")
}
