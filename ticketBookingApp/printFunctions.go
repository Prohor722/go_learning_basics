package main

import "fmt"

func showMoviesList() {
	fmt.Println("Available Movies:")
	for i, movie := range movies {
		fmt.Printf("%d. %s ($%.2f) - %d tickets available\n", i+1, movie.Name, movie.Price, movie.AvailableTickets)
	}
}