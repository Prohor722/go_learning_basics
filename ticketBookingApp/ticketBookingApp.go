package main

import "fmt"

type BoughtTicket struct {
	id          int
	Name        string
	email       string
	noOfTickets int
}

// var boughtTickets []BoughtTicket // Removed duplicate declaration

func ticketBookingApp() {
	var name string
	var age int
	var email string
	var bookTickets int
	

	name = getName()
	age = getAge()
	email = getEmail()

	showMoviesList()

	fmt.Print("\nEnter the movie number:")
	var movieChoice int
	movieChoice = getChoice()

	selectedMovie := movies[movieChoice-1]
	fmt.Printf("You have selected %s. There are %d tickets available. Each ticket costs $%.2f.\n", selectedMovie.Name, selectedMovie.AvailableTickets, selectedMovie.Price)

	bookTickets = getNoOfTickets()


	if age < 18 {
		fmt.Println("Sorry, you must be at least 18 years old to book tickets.")
		return
	}
	if selectedMovie.AvailableTickets <= 0 {
		fmt.Println("Sorry, no tickets are available.")
		return
	}
	if bookTickets > selectedMovie.AvailableTickets {
		fmt.Printf("Sorry, only %d tickets are available.\n", selectedMovie.AvailableTickets)
		return
	}else if bookTickets <= 0 {
		fmt.Println("Please enter a valid number of tickets to book.")
		return
	}

	totalCost := float64(bookTickets) * selectedMovie.Price
	ticketsAvailable := selectedMovie.AvailableTickets
	ticketsAvailable -= bookTickets
	
	boughtTickets = append(boughtTickets, BoughtTicket{
		id: selectedMovie.id,
		Name: name,
		email: email,
		noOfTickets: bookTickets,
	})

	fmt.Printf("Thank you %s for booking %d tickets. Total cost: $%.2f\n", name, bookTickets, totalCost)
	fmt.Printf("Against %s this email address.\n", email)
	fmt.Printf("Tickets remaining: %d\n", ticketsAvailable)	
}
