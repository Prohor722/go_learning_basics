package main

import "fmt"

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


	if(!buyTicketValidation(age, selectedMovie.AvailableTickets, bookTickets, selectedMovie.ageLimit)) {
		return
	}

	totalCost := float64(bookTickets) * selectedMovie.Price
	ticketsAvailable := selectedMovie.AvailableTickets
	ticketsAvailable -= bookTickets
	
	boughtTickets = append(boughtTickets, BoughtTickets{
		id: selectedMovie.id,
		Name: name,
		email: email,
		noOfTickets: bookTickets,
	})

	fmt.Printf("Thank you %s for booking %d tickets. Total cost: $%.2f\n", name, bookTickets, totalCost)
	fmt.Printf("Against %s this email address.\n", email)
	fmt.Printf("Tickets remaining: %d\n", ticketsAvailable)	
}
