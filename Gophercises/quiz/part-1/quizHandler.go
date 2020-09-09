package main

import "fmt"

var noOfCorrectAnswers int
var totalQuestions int

func takeQuiz(problemSet []Problem) int {
	noOfCorrectAnswers = 0
	totalQuestions = len(problemSet)

	for i, prob := range problemSet{
		fmt.Print("Problem #", i, ": ", prob.Question, " = " )
		var input int
		fmt.Scanf("%d", &input)
		if input == prob.Answer {
			noOfCorrectAnswers++
		}
	}

	return noOfCorrectAnswers
}

func result() {
	fmt.Println("You scored ", noOfCorrectAnswers, " out of ", totalQuestions)
}