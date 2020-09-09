package main

import (
	"strconv"
	"strings"
)

type Problem struct {
	Question string
	Answer int
}

func buildProblemSetFromCSV(csvLines [][]string) []Problem {
	problemSet := make([]Problem,0)
	for _, line := range csvLines {

		ans, _ := strconv.Atoi(strings.TrimSpace(line[1]))

		problemSet = append(problemSet, Problem{
			Question: line[0],
			Answer:	ans,
		})
	}
	return problemSet
}