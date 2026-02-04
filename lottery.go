package main

import (
	"fmt"
	"math/rand")

type Tickets struct {
	Number string
	Name string
	Phone string
}
var tickets []*Tickets
var lastSoldTicketIndex int = -1

func drawLottery() {
	var numberOfTickets int
	var winningTicketIndex int
	fmt.Print("Welcome To the Lottery Game !!\n")
	fmt.Printf("Number of Tickets to play the Game:")
	fmt.Scan(&numberOfTickets)
	generateTickets(numberOfTickets)
	buyTickets()
	winningTicketIndex = winningNumber()

	if( winningTicketIndex < 1 ) {
		fmt.Println("Not enough tickets sold to draw a winner.")
	}else{
		fmt.Printf("The winning ticket number is: %d\n", tickets[winningTicketIndex-1].Number)
	}



	// fmt.Printf("Winning Numbers: %v\n", winningNumbers)
}
func winningNumber() int {
	if len(tickets) < 2 {
		fmt.Println("Not enough tickets sold to draw a winner.")
		return 0
	}
	return rand.Intn(len(tickets))
}

func generateTickets(numberOfTickets int) {
	fmt.Println("Generating lottery tickets...")
	for i := 0; i < numberOfTickets; i++ {
		ticket := Tickets{Number: string{i+""+i+1+i + 2, i + 3, i + 4}, Name: "", Phone: ""}
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
