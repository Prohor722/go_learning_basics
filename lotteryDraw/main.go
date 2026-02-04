package main

import (
	"fmt"
)

type Tickets struct {
	Number string
	Name string
	Phone string
}
var tickets []*Tickets
var lastSoldTicketIndex int = -1
var numberOfTickets int

func main(){
	fmt.Printf("Number of Tickets to play the Game:")
	fmt.Scan(&numberOfTickets)
	validNumber(numberOfTickets)
	
	fmt.Print("Welcome To the Lottery Game !!\n")
	drawLottery()
}