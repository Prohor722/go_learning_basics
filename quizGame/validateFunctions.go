package main

import (
	"strings"
)

func validateName (name string) bool {
	name = strings.TrimSpace(name)
	if len(name) < 2 {
		println("Please enter a valid name (at least 2 characters).")
		return false
	}
	return true
}

func validateAge(age uint) bool {
	if age <= 1 {
		println("Invalid age entered. Please enter a valid age.")
		return false
	}
	return true
}