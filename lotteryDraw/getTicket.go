package main

import "fmt"

func getTicket(buyTickets int, name string, phone string) bool {
	var logic1 bool = lastSoldTicketIndex >= len(tickets)-1
	var logic2 bool = tickets[lastSoldTicketIndex+1].Name != ""
	var logic3 bool = lastSoldTicketIndex+buyTickets > totalLotteryTickets

	for i := 0; i < buyTickets; i++ {
		if logic1 || logic2 || logic3 {
			// fmt.Printf("lastSoldTicketIndex >= len(tickets)-1: %v",lastSoldTicketIndex >= len(tickets)-1)
			// fmt.Printf("tickets[lastSoldTicketIndex+1].Name != \"\": %v", tickets[lastSoldTicketIndex+1].Name != "")
			fmt.Printf("lastSoldTicketIndex+buyTickets > totalLotteryTickets: %v\n", logic3)
			fmt.Printf("lastSoldTicketIndex: %d\n",lastSoldTicketIndex)
			fmt.Printf("buyTickets: %d\n", buyTickets)
			fmt.Printf("totalLotteryTickets: %d\n", totalLotteryTickets)
			return false
		}
		lastSoldTicketIndex++
		tickets[lastSoldTicketIndex].Name = name
		tickets[lastSoldTicketIndex].Phone = phone
	}
	return true
}
