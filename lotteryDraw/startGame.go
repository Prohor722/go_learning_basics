package main

import "fmt"

func startGame() {
	fmt.Print("Welcome To the Lottery Game !!\n")
	fmt.Printf("Number of Tickets to play the Game:")
	fmt.Scan(&totalLotteryTickets)
	
	if(!validNumber(totalLotteryTickets) && totalLotteryTickets < 2){
		fmt.Println("Invalid number of tickets. Exiting the game. (Min:2)")
		return
	}
	if(!generateTickets(totalLotteryTickets)) {
		fmt.Println("Failed to generate tickets. Restarting the game.")
		startGame()
		return
	}
	fmt.Println()
	drawLottery()
}