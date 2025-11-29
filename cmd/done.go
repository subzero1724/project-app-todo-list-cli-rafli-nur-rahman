package cmd

import (
	"fmt"
	"project-app-todo-list-cli-rafli-nur-rahman/service"
	"strconv"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [id]",
	Short: "Mark a task as done",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := strconv.Atoi(args[0])

		err := service.MarkTaskDone(id)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Task marked as done")
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
