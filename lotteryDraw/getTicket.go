package main

import "fmt"

func getTicket(buyTickets int, name string, phone string) bool {

	for i := 0; i < buyTickets; i++ {
		if lastSoldTicketIndex >= len(tickets)-1 || tickets[lastSoldTicketIndex+1].Name != "" || lastSoldTicketIndex+1+buyTickets > totalLotteryTickets {
			fmt.Printf("lastSoldTicketIndex >= len(tickets)-1: %v",lastSoldTicketIndex >= len(tickets)-1)
			fmt.Printf("tickets[lastSoldTicketIndex+1].Name != \"\": %v", tickets[lastSoldTicketIndex+1].Name != "")
			fmt.Printf("lastSoldTicketIndex+1+buyTickets > totalLotteryTickets: %v", lastSoldTicketIndex+1+buyTickets > totalLotteryTickets)
			return false
		}
		lastSoldTicketIndex++
		tickets[lastSoldTicketIndex].Name = name
		tickets[lastSoldTicketIndex].Phone = phone
	}
	return true
}
