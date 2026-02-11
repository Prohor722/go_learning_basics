package main

import (
	"fmt"
	"strings"
)


func quizGame() {
	var name string
	var age uint = 0

	println("Welcome to the Quiz Game!")
	name = getName()
	age = getAge(name)

	println("Hello", name, "age", age, "! Let's start the quiz.")

	score := 0
	
	println("Quiz Over! Your total score is:", score, "out of", len(questions))
	fmt.Printf("Your answer is %v%% correct!", (score*100)/len(questions))
	println("\nThank you for playing, ", name, "!")

}

func getUserAnswer(questions []string, options [][]string) int {
	var score int = 0
	for i, question := range questions {
		println(question)
		for _, option := range options[i] {
			println(option)
		}
		var answer string
		print("Your answer: ")
		fmt.Scanln(&answer)

		if strings.ToUpper(answer) == correctAnswers[i] {
			println("Correct!")
			score++
		} else {
			println("Wrong! The correct answer is", correctAnswers[i])
		}
	}
}