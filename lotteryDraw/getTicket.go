package main

func getTicket(numberOfTickets int, name string, phone string) bool {

	for i := 0; i < numberOfTickets; i++ {
		if( lastSoldTicketIndex+1 >= len(tickets) || tickets[lastSoldTicketIndex+1].Name != "" || lastSoldTicketIndex+1+numberOfTickets > len(tickets) ){ 
			return false
		}
		lastSoldTicketIndex++
		tickets[lastSoldTicketIndex].Name = name
		tickets[lastSoldTicketIndex].Phone = phone
	}
	return true
}
