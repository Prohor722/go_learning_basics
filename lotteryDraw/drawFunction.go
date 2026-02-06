package main

import (
	"fmt"
)

func drawLottery() {
	var winningTicketIndex int

	if(lastSoldTicketIndex < len(tickets)-1){
		drawLottery()
	}

	winningTicketIndex = winningNumber()

	if( winningTicketIndex < 1 ) {
		fmt.Println("Not enough tickets sold to draw a winner.")
	}else{
		fmt.Printf("The winning ticket number is: %s\n Winner Name: %s & Phone: %s\n", tickets[winningTicketIndex-1].Number, tickets[winningTicketIndex-1].Name, tickets[winningTicketIndex-1].Phone)
	}
}