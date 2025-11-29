package cmd

import (
	"fmt"
	"project-app-todo-list-cli-rafli-nur-rahman/service"

	"github.com/spf13/cobra"
)

var desc string

var addCmd = &cobra.Command{
	Use:   "add [title]",
	Short: "Add a new task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]

		err := service.AddTask(title, desc)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Task added successfully!")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&desc, "desc", "d", "", "Task description")
}
