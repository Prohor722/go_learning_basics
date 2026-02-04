package main

import (
	"fmt"
)

func buyTickets() bool {
	var buyTickets int
	var name string
	var phone string
	fmt.Print("\nEnter the name of the ticket buyer:")
	fmt.Scan(&name)
	if !validateName(name) {
		return false
	}	
	fmt.Print("\nEnter the phone number of the ticket buyer:")
	fmt.Scan(&phone)
	if !validatePhone(phone) {
		return false
	}
	fmt.Print("\nHow many tickets would you like to buy?")
	fmt.Scan(&buyTickets)
	if !validNumber(buyTickets) {
		return false
	}

	if( getTicket(buyTickets,name,phone) != false ) {
		fmt.Printf("You have successfully bought ticket with name: %s and phone: %s\n", name, phone)
	}else{
		fmt.Println("Sorry, not enough tickets available.")
	}
	return true
}
