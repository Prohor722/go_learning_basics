package main

import (
	"fmt"
)

func drawLottery() {
	var winningTicketIndex int
	fmt.Print("Welcome To the Lottery Game !!\n")
	fmt.Printf("Number of Tickets to play the Game:")
	fmt.Scan(&numberOfTickets)
	validateFunctionReturns(generateTickets(numberOfTickets))
	validateFunctionReturns(buyTickets())
	winningTicketIndex = winningNumber()

	if( winningTicketIndex < 1 ) {
		fmt.Println("Not enough tickets sold to draw a winner.")
	}else{
		fmt.Printf("The winning ticket number is: %s\n", tickets[winningTicketIndex-1].Number)
	}




	// fmt.Printf("Winning Numbers: %v\n", winningNumbers)
}