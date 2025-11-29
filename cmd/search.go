package cmd

import (
	"project-app-todo-list-cli-rafli-nur-rahman/service"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [keyword]",
	Short: "Search tasks by keyword",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		keyword := args[0]           // <-- fix: ambil dari argumen
		service.SearchTasks(keyword) // <-- fix: pakai variabel
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
