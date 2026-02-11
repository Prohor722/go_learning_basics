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

	score := getUserAnswer(questionsA, optionsA, correctAnswersA)
	
	println("\nThank you for playing, ", name, "!")

}
