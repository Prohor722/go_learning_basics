package main

import (
	"fmt"
	"math"
	"strings"
)

func getName() string {
	var name string
	print("\nEnter your name: ")
	fmt.Scanln(&name)
	if !validateName(name) {
		return getName()
	}
	return name
}

func getAge(name string) uint {
	var age uint = 0
	print("Enter your age: ")
	fmt.Scanln(&age)
	age = uint(math.Abs(float64(age)))
	if !validateAge(age) {
		return getAge(name)
	}
	return age
}

func getUserAnswer(questions []Question) {
	var score int = 0
	for _, question := range questions {
		println(question.question)
		for _, option := range question.options {
			println(option)
		}
		var answer string
		print("Your answer: ")
		fmt.Scanln(&answer)

		if strings.ToUpper(answer) == question.correctAnswer {
			println("Correct!")
			score++
		} else {
			println("Wrong! The correct answer is", question.correctAnswer)
		}
	}
	println("Quiz Over! Your total score is:", score, "out of", len(questions))
	fmt.Printf("Your answer is %v%% correct!", (score*100)/len(questions))
}