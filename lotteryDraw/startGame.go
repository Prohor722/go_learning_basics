package main

import "fmt"

func startGame() {
	fmt.Printf("Number of Tickets to play the Game:")
	fmt.Scan(&totalLotteryTickets)
	
	if(!validNumber(totalLotteryTickets)){
		fmt.Println("Invalid number of tickets. Exiting the game.")
		return
	}
	if(!generateTickets(totalLotteryTickets)) {
		fmt.Println("Failed to generate tickets. Restarting the game.")
		startGame()
		return
	}

	fmt.Print("Welcome To the Lottery Game !!\n")
	drawLottery()
}