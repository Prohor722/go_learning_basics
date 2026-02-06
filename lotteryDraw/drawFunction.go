package main

import (
	"fmt"
)

func drawLottery() bool {
	if(lastSoldTicketIndex < len(tickets)-1){
		fmt.Printf("Sorry all tickets are not sold yet!! \n")
		availableTickets()
		return false
	}

	winningTicketIndex := winningNumber()
	winningTicket := tickets[winningTicketIndex-1]
	winnerName := winningTicket.Name
	winnerPhone := winningTicket.Phone

	if( len(tickets) < 2) {
		fmt.Println("Not enough tickets sold to draw a winner.")
	}else{
		fmt.Printf("The winning ticket number is: %s\n Winner Name: %s & Phone: %s\n", winningTicket.Number, winnerName, winnerPhone)
		return true
	}

	return false
}