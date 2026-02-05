package main

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
