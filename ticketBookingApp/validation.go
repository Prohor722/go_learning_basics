package main

import (
	"fmt"
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

func validateMovieChoice(choice int) bool {
	if choice < 1 || choice > len(movies) {
		return false
	}
	return true
}

func buyTicketValidation(age,availableTickets, bookTickets int) bool {
	if age < 18 {
		fmt.Println("Sorry, you must be at least 18 years old to book tickets.")
		return false
	}
	if availableTickets <= 0 {
		fmt.Println("Sorry, no tickets are available.")
		return false
	}
	if bookTickets > availableTickets {
		fmt.Printf("Sorry, only %d tickets are available.\n", availableTickets)
		return false
	}else if bookTickets <= 0 {
		fmt.Println("Please enter a valid number of tickets to book.")
		return false
	}
	return true
}