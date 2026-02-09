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

	fmt.Printf("Which movie do you want to watch?")
	for i, movie := range movies {
		fmt.Printf("\n%d. %s ($%.2f) - %d tickets available", i+1, movie.Name, movie.Price, movie.AvailableTickets)
	}
	fmt.Print("\nEnter the movie number:")
	var movieChoice int
	fmt.Scanln(&movieChoice)
	if movieChoice < 1 || movieChoice > len(movies) {
		fmt.Println("Invalid movie choice. Please try again.")
		return
	}
	selectedMovie := movies[movieChoice-1]
	fmt.Printf("You have selected %s. There are %d tickets available. Each ticket costs $%.2f.\n", selectedMovie.Name, selectedMovie.AvailableTickets, selectedMovie.Price)

	fmt.Print("Enter number of tickets to book:")
	fmt.Scanln(&bookTickets)


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
	fmt.Printf("Thank you %s for booking %d tickets. Total cost: $%.2f\n", name, bookTickets, totalCost)
	fmt.Printf("Tickets remaining: %d\n", ticketsAvailable)	
}
