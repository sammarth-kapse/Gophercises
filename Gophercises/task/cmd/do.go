package cmd

import (
	"fmt"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Completes the given task.",
	Long: "Completes the given task.\nRequires exactly 1 argument.\nThe task which has to be completed.",
	Args: cobra.RangeArgs(1,1),
	Run: func(cmd *cobra.Command, args []string) {
		taskNum, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task number.")
		}
		checkError(err)

		tasks, _ := db.GetTaskList()
		if taskNum > len(tasks) {
			fmt.Println("Invalid Task number")
			return
		}

		taskID := tasks[taskNum-1].Key
		taskValue := tasks[taskNum-1].Value

		err = db.DeleteTask(taskID)
		if err != nil {
			fmt.Println("Cannot List tasks. Error Occurred.\nError: ", err)
		}
		fmt.Printf("You have completed \"%s\" task.\n", taskValue)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
