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