package main

import (
	"strings"
)

func nameValidation(name string) bool {
	if len(name) < 2 {
		return false
	}
	return true
}

func ageValidation(age int) bool {
	if age < 0 || age > 120 {
		return false
	}
	return true
}

func emailValidation(email string) bool {
	if len(email) < 5 || !strings.Contains(email, "@") {
		return false
	}
	return true
}

func noOfTicketsValidation(noOfTickets int) bool {
	if noOfTickets <= 0 {
		return false
	}
	return true
}