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
	startGame()
}

func startGame() {
	fmt.Printf("Number of Tickets to play the Game:")
	fmt.Scan(&numberOfTickets)
	
	if(!validNumber(numberOfTickets)){
		fmt.Println("Invalid number of tickets. Exiting the game.")
		return
	}
	if(!generateTickets(numberOfTickets)) {
		fmt.Println("Failed to generate tickets. Restarting the game.")
		startGame()
		return
	}

	fmt.Print("Welcome To the Lottery Game !!\n")
	drawLottery()
}