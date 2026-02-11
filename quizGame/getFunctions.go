package main

import (
	"fmt"
	"math"
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

func getUserAnswer(questions []string, options [][]string, correctAnswers []string) int {
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

	return score
}