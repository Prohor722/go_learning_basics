package main

import "fmt"

func ticketBookingApp() {
	var name string
	var age int
	var email string
	var ticketsAvailable int = 100
	var ticketPrice float64 = 50.0
	var bookTickets int

	fmt.Print("Enter your name:")
	fmt.Scanln(&name)

	fmt.Print("Enter your age:")
	fmt.Scanln(&age)

	fmt.Print("Enter your email:")
	fmt.Scanln(&email)

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

	




}