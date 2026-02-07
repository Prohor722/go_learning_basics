package main

import "fmt"

type Tickets struct {
	Number string
	Name   string
	Phone  string
}

var tickets []*Tickets
var lastSoldTicketIndex int = -1
var totalLotteryTickets int
var winningTicketIndex int

func main() {
	fmt.Print("Welcome To the Lottery Game !!\n")
	fmt.Printf("Number of Tickets to play the Game:")
	totalLotteryTickets = getNumber()

	if !generateTickets(totalLotteryTickets) {
		fmt.Println("Failed to generate tickets. Restarting the game.")
		fmt.Println()
		main()
		return
	}
	startGame()
}
