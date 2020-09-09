package main

import (
	"flag"
	"fmt"
)

func main(){
	csvFileName := flag.String("csv", "problem.csv", "CSV file")
	quizTime := flag.Int("limit", 3, "timer")
	flag.Parse()

	askToStart()
	csvLines := readCsvLineWise(*csvFileName)
	problemSet := buildProblemSetFromCSV(csvLines)
	takeQuiz(problemSet, *quizTime)
	result()
}

func askToStart() {
	var temp int
	fmt.Println("Press any Key to start")
	fmt.Scanf("%d", &temp)
	fmt.Println()
}