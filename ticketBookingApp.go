package main

import (
	"fmt"
	"strings"
)

func ticketBookingApp() {
	var name string
	var age int
	var email string
	var ticketsAvailable int = 100
	var ticketPrice float64 = 50.0
	var availableMoviesAndPrices = map[string]float64{
		"Movie A": 50.0,
		"Movie B": 60.0,
		"Movie C": 70.0,
	}
	var bookTickets int

	fmt.Print("Enter your name:")
	fmt.Scanln(&name)

	if(!validation("name", name)) {
		fmt.Println("Invalid name. Please try again.")
		return
	}

	fmt.Print("Enter your age:")
	fmt.Scanln(&age)

	if(!validation("age", age)) {
		fmt.Println("Invalid age. Please enter a valid age between 0 and 120.")
		return
	}

	fmt.Print("Enter your email:")
	fmt.Scanln(&email)
	if(!validation("email", email)) {
		fmt.Println("Invalid email. Please enter a valid email address.")
		return
	}

	fmt.Printf("Welcome %s! There are %d tickets available at $%.2f each.\n", name, ticketsAvailable, ticketPrice)

	fmt.Println("Available Movies and Prices:")
	for movie, price := range availableMoviesAndPrices {
		fmt.Printf("%s: $%.2f\n", movie, price)
	}

	fmt.Printf("Select Movie:")
	// For simplicity, we are not validating movie selection here
	var selectedMovie string
	fmt.Scanln(&selectedMovie)
	fmt.Print("Enter number of tickets to book:")
	fmt.Scanln(&bookTickets)


	if age < 18 {
		fmt.Println("Sorry, you must be at least 18 years old to book tickets.")
		return
	}
	if ticketsAvailable <= 0 {
		fmt.Println("Sorry, no tickets are available.")
		return
	}
	if bookTickets > ticketsAvailable {
		fmt.Printf("Sorry, only %d tickets are available.\n", ticketsAvailable)
		return
	}else if bookTickets <= 0 {
		fmt.Println("Please enter a valid number of tickets to book.")
		return
	}

	totalCost := float64(bookTickets) * ticketPrice
	ticketsAvailable -= bookTickets
	fmt.Printf("Thank you %s for booking %d tickets. Total cost: $%.2f\n", name, bookTickets, totalCost)
	fmt.Printf("Tickets remaining: %d\n", ticketsAvailable)
}

func validation(validationType string, value interface{}) bool {
	switch validationType {
	case "name":
		name, ok := value.(string)
		if !ok || len(name) < 2 {
			return false
		}
		return true
	case "age":
		age, ok := value.(int)
		if !ok || age < 0 || age > 120 {
			return false
		}
		return true
	case "email":
		email, ok := value.(string)
		if !ok || len(email) < 5 || !strings.Contains(email, "@") {
			return false
		}
		return true
	default:
		return false
	}
}