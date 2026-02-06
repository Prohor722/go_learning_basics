package main

import (
	"fmt"
)

func startGame() {

	fmt.Printf("Select a option:\n")
	fmt.Printf("1. Buy Tickets\n")
	fmt.Printf("2. Show Number of Available Tickets\n")
	fmt.Printf("3. Check your Tickets\n")
	fmt.Printf("4. Run Draw\n")
	fmt.Printf("5. Exit")
	fmt.Print("Enter your choice:")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		if(!buyTickets()){
			fmt.Println("Ticket purchase failed. Restarting ticket purchase.")
		}
	case 2:
		availableTickets()
	case 3:
		fmt.Printf("Enter your phone number to check your tickets:")
		showUserTickets(getName(), getPhone())
	case 4:
		drawLottery()
	case 5:
		fmt.Println("Thank you for playing the Lottery Game. Goodbye!")
		return
	default:
		fmt.Println("Invalid choice. Restarting the game.")
	}
	
	// if(!validNumber(totalLotteryTickets) || totalLotteryTickets < 2){
	// 	fmt.Println("Invalid number of tickets. Exiting the game. (Min:2)")
	// 	return
	// }
	// if(!generateTickets(totalLotteryTickets)) {
	// 	fmt.Println("Failed to generate tickets. Restarting the game.")
	// 	startGame()
	// 	return
	// }
	// if(!buyTickets()){
	// 	fmt.Println("Ticket purchase failed. Restarting ticket purchase.")
	// 	drawLottery()
	// 	return
	// }
	// fmt.Println()
	startGame()
}