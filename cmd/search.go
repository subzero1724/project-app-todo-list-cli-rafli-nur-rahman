package cmd

import (
	"fmt"
	"strings"

	"project-app-todo-list-cli-rafli-nur-rahman/model"
	"project-app-todo-list-cli-rafli-nur-rahman/service"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [keyword]",
	Short: "Search tasks by keyword",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a keyword to search!")
			return
		}

		keyword := strings.Join(args, " ")

		allTasks := service.GetAllTasks()
		var results []model.Task

		for _, t := range allTasks {
			if strings.Contains(strings.ToLower(t.Title), strings.ToLower(keyword)) ||
				strings.Contains(strings.ToLower(t.Description), strings.ToLower(keyword)) {
				results = append(results, t)
			}
		}

		if len(results) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		displayTasks(results)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
