package cmd

import (
	"fmt"
	"project-app-todo-list-cli-rafli-nur-rahman/service"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := strconv.Atoi(args[0])

		err := service.DeleteTask(id)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Task deleted successfully!")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
