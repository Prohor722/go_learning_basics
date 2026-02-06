package main

import (
	"fmt"
)

func drawLottery() bool {
	var winningTicketIndex int
	winningTicket := tickets[winningTicketIndex-1]

	if(lastSoldTicketIndex < len(tickets)-1){
		drawLottery()
	}

	winningTicketIndex = winningNumber()

	if( winningTicketIndex < 1 ) {
		fmt.Println("Not enough tickets sold to draw a winner.")
	}else{
		fmt.Printf("The winning ticket number is: %s\n Winner Name: %s & Phone: %s\n", winningTicket.Number, winningTicket.Name, winningTicket.Phone)
		return true
	}

	return false
}