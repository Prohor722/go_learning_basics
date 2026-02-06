package main

import "fmt"

func getTicket(buyTickets int, name string, phone string) bool {
	var logic1 bool = lastSoldTicketIndex >= len(tickets)-1
	var logic2 bool = tickets[lastSoldTicketIndex+1].Name != ""
	var logic3 bool = lastSoldTicketIndex+buyTickets > totalLotteryTickets

	for i := 0; i < buyTickets; i++ {
		if logic1 || logic2 || logic3 {
			return false
		}
		lastSoldTicketIndex++
		tickets[lastSoldTicketIndex].Name = name
		tickets[lastSoldTicketIndex].Phone = phone
	}
	return true
}

func availableTickets() {
	availableTickets := totalLotteryTickets - (lastSoldTicketIndex + 1)

	if( availableTickets < 1 ) {
		fmt.Println("No tickets available!! You can run draw to check the Winner.")
		return
	}
	fmt.Printf("Currently available tickets: %d\n", availableTickets)
}

func showUserTickets(name string, phone string) {
	var userTickets []string
	for i := 0; i <= lastSoldTicketIndex; i++ {
		if tickets[i].Name == name && tickets[i].Phone == phone {
			userTickets = append(userTickets, tickets[i].Number)
		}
	}
	if len(userTickets) == 0 {
		println("No tickets found for the provided name and phone number.")
	} else {
		fmt.Printf("You have bought %v tickets.\n", len(userTickets))
			println("Your Tickets:")
		for _, ticket := range userTickets {
			println(ticket)
		}
	}
}