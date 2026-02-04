package main

import (
	"fmt"
)

func drawLottery() {
	var winningTicketIndex int
	validateFunctionReturns(generateTickets(numberOfTickets))
	validateFunctionReturns(buyTickets())

	if(lastSoldTicketIndex < len(tickets)-1){
		drawLottery()
	}

	winningTicketIndex = winningNumber()

	if( winningTicketIndex < 1 ) {
		fmt.Println("Not enough tickets sold to draw a winner.")
	}else{
		fmt.Printf("The winning ticket number is: %s\n", tickets[winningTicketIndex-1].Number)
	}
}