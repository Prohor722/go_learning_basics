package main

import "fmt"

type Tickets struct {
	Numbers []int
	Name string
	Phone string
}
var tickets []*Tickets
var lastSoldTicketIndex int = -1

func drawLottery(winningNumbers []int, userNumbers []int, contact *Contact) {
	fmt.Print("Welcome To the Lottery Game !!\n")
	fmt.Printf("Winning Numbers: %v\n", winningNumbers)
}

func generateTickets(numberOfTickets int) {
	fmt.Println("Generating lottery tickets...")
	for i := 0; i < numberOfTickets; i++ {
		ticket := Tickets{Numbers: []int{i, i + 1, i + 2, i + 3, i + 4}, Name: "", Phone: ""}
		tickets = append(tickets, &ticket)
	}
}

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

func buyTickets(){
	var buyTickets int
	var name string
	var phone string
	fmt.Print("\nEnter the name of the ticket buyer:")
	fmt.Scan(&name)
	if !validateName(name) {
		return
	}	
	fmt.Print("\nEnter the phone number of the ticket buyer:")
	fmt.Scan(&phone)
	if !validatePhone(phone) {
		return
	}
	fmt.Print("\nHow many tickets would you like to buy?")
	fmt.Scan(&buyTickets)

	if( getTicket(buyTickets,name,phone) != false ) {
		fmt.Printf("You have successfully bought ticket with name: %s and phone: %s\n", name, phone)
	}else{
		fmt.Println("Sorry, not enough tickets available.")
	}
}

func validateName(name string) bool {
	var minLength = 2
	if len(name) < minLength {
		fmt.Println("Name is too short.(Min:", minLength, ")")
		return false
	}
	return true
}
func validatePhone(phone string) bool {
	var minLength = 5
	if len(phone) < minLength {
		fmt.Printf("Phone number is too short.(Min: %d)\n", minLength)
		return false
	}
	return true
}
