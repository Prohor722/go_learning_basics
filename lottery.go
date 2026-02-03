package main

import "fmt"

type Tickets struct {
	Numbers []int
}

tickets := []Tickets{}
var lastSoldTicketIndex int = -1

func drawLottery(winningNumbers []int, userNumbers []int, contact *Contact) {

}

func generateTickets(){
	fmt.Println("Generating lottery tickets...")
	for i := 0; i < 5; i++ {
		ticket := Tickets{Numbers: []int{i, i + 1, i + 2, i + 3, i + 4}}
		tickets = append(tickets, ticket)
	}
}

func getTicket(numberOfTickets int) *Tickets {
	if( lastSoldTicketIndex+1 >= len(tickets) ) {
		return nil
	}
	lastSoldTicketIndex += numberOfTickets
	return &tickets[lastSoldTicketIndex]
}

func buyTickets(){
	var buyTickets int
	fmt.Print("How many tickets would you like to buy?")
	fmt.Scan(&buyTickets)
}

