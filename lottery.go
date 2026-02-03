package main

import "fmt"

type Tickets struct {
	Numbers []int
	Owner string
}

tickets := []Tickets{}
var lastSoldTicketIndex int = -1

func drawLottery(winningNumbers []int, userNumbers []int, contact *Contact) {

}

func generateTickets(){
	fmt.Println("Generating lottery tickets...")
	for i := 0; i < 5; i++ {
		ticket := Tickets{Numbers: []int{i, i + 1, i + 2, i + 3, i + 4}, Owner: ""}
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
	var name string
	var phone string
	fmt.Print("Enter the name of the ticket buyer:")
	fmt.Scan(&name)
	fmt.Print("\nEnter the phone number of the ticket buyer:")
	fmt.Scan(&phone)
	fmt.Print("\nHow many tickets would you like to buy?")
	fmt.Scan(&buyTickets)

	if( ticket := getTicket(buyTickets); ticket != nil ) {
		fmt.Printf("You have successfully bought ticket with numbers: %v\n", ticket.Numbers)
	}else{
		fmt.Println("Sorry, not enough tickets available.")
	}
}

