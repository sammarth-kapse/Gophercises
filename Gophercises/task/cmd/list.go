package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"task/db"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the task.",
	Run: func(cmd *cobra.Command, args []string) {

		tasks, err := db.GetTaskList()
		if err != nil {
			fmt.Println("Cannot List tasks. Error Occurred.\nError: ", err)
		}

		if len(tasks) > 0 {
			fmt.Println("You have the following tasks:")
		} else {
			fmt.Println("There are no pending tasks.")
		}

		for i, t := range tasks {
			fmt.Println(i+1, ". ", t.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
