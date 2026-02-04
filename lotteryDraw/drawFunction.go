package main

import (
	"fmt"
)

func drawLottery() {
	var winningTicketIndex int
	if(!buyTickets()){
		fmt.Println("Ticket purchase failed. Restarting ticket purchase.")
		drawLottery()
		return
	}

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