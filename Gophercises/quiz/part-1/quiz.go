package main

import "flag"

func main(){
	csvFileName := flag.String("csv", "problem.csv", "CSV file")
	flag.Parse()

	csvLines := readCsvLineWise(*csvFileName)
	problemSet := buildProblemSetFromCSV(csvLines)
	takeQuiz(problemSet)
	result()
}
