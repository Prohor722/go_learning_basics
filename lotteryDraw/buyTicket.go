package main

import (
	"fmt"
)

func buyTickets() bool {
	var buyTickets int
	var name string
	var phone string
	fmt.Print("\nEnter the name of the ticket buyer:")
	name = getName()
	fmt.Print("\nEnter the phone number of the ticket buyer:")
	phone = getPhone()
	fmt.Print("\nHow many tickets would you like to buy:")
	fmt.Scan(&buyTickets)
	if !validNumber(buyTickets) {
		return false
	}

	if( getTicket(buyTickets,name,phone) != false ) {
		fmt.Printf("You have successfully bought ticket with name: %s and phone: %s\n", name, phone)
		fmt.Printf("Your Ticket numbers are:")
		for i := lastSoldTicketIndex - buyTickets + 1; i <= lastSoldTicketIndex; i++ {
			fmt.Printf(" %s ", tickets[i].Number)
		}
		fmt.Println()
	}else{
		fmt.Println("Sorry, not enough tickets available.")
		return false
	}
	return true
}
