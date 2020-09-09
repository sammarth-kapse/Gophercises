package main

import (
	"task/cmd"
	"task/db"
)

func main() {
	err :=db.Init("tasks")
	if err != nil {
		panic(err)
	}
	cmd.RootCmd.Execute()
}


