package main

import (
	"fmt"
	"time"
)

var noOfCorrectAnswers int
var totalQuestions int

func takeQuiz(problemSet []Problem, quizTime int) {
	noOfCorrectAnswers = 0
	totalQuestions = len(problemSet)

	timer := time.NewTimer(time.Duration(quizTime)*time.Second)
	quizChannel := make(chan int)
	go askQuestions(problemSet, quizChannel)
	select {
	case <-timer.C:
		return
	case <-quizChannel:
		break
	}

}

func askQuestions(problemSet []Problem, quizChannel chan int) {
	for i, prob := range problemSet{
		fmt.Print("Problem #", i+1, ": ", prob.Question, " = " )
		var input int
		fmt.Scanf("%d", &input)
		if input == prob.Answer {
			noOfCorrectAnswers++
		}
	}
	quizChannel <- noOfCorrectAnswers
}


func result() {
	fmt.Println("\nYou scored ", noOfCorrectAnswers, " out of ", totalQuestions)
}