package main

import (
	"fmt"
	"strings"
)

func quizGame() {
	var name string
	var age int

	println("Welcome to the Quiz Game!")
	print("Enter your name: ")
	fmt.Scanln(&name)
	print("Enter your age: ")
	fmt.Scanln(&age)

	println("Hello", name, "age", age, "! Let's start the quiz.")

	questions := []string{
		"What is the capital of France?",
		"What is 2 + 2?",
		"What is the largest planet in our solar system?",
	}

	options := [][]string{
		{"A) Berlin", "B) Madrid", "C) Paris", "D) Rome"},
		{"A) 3", "B) 4", "C) 5", "D) 6"},
		{"A) Earth", "B) Jupiter", "C) Mars", "D) Saturn"},
	}
	correctAnswers := []string{"C", "B", "B"}
	score := 0
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
	println("Quiz Over! Your total score is:", score, "out of", len(questions))

}
