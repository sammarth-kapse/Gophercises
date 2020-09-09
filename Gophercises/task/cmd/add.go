
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	"task/db"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds the given task to task-list.",
	Long: "Adds the given task to task-list.\nRequires at least 1 argument.\nConcatenation of all the argument is the task that will be added.",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		newTask := strings.Join(args, " ")

		err := db.AddTask(newTask)
		if err != nil {
			fmt.Println("Cannot List tasks. Error Occurred.\nError: ", err)
		}

		fmt.Printf("Added \"%s\" to your task list.\n", newTask)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)

}
