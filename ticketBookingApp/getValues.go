package main

import "fmt"

func getName() []string {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	if !nameValidation(name) {
		fmt.Println("Invalid name. Please try again.")
		return getName()
	}
	return []string{name}
}

func getAge() []int {
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)
	if !ageValidation(age) {
		fmt.Println("Invalid age. Please enter a valid age between 0 and 120.")
		return getAge()
	}
	return []int{age}
}