package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var RootCmd = &cobra.Command{
	Use: "task",
	Short: "Task Manager for you!",
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
	return
}