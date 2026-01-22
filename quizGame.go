package main

import "fmt"

func quizGame(){
	var name string
	var age int

	println("Welcome to the Quiz Game!")
	print("Enter your name: ")
	fmt.Scanln(&name)
	print("Enter your age: ")
	fmt.Scanln(&age)

	println("Hello", name, "age", age, "! Let's start the quiz.")
}