package main

import "fmt"

func ticketBookingApp() {
	var name string
	var age int
	var ticketsAvailable int = 100

	fmt.Print("Enter your name:")
	fmt.Scanln(&name)

	fmt.Print("Enter your age:")
	fmt.Scanln(&age)
}