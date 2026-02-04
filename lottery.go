package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tickets struct {
	Number string
	Name string
	Phone string
}
var tickets []*Tickets
var lastSoldTicketIndex int = -1
var numberOfTickets int

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomSixDigitString(id int) string {
	n := rand.Intn(1000000)+id // 0..999999
	if( n > 999999 ) {
		n -= numberOfTickets*2
	}
	return fmt.Sprintf("%06d", n)
}


func winningNumber() int {
	if len(tickets) < 2 {
		fmt.Println("Not enough tickets sold to draw a winner.")
		return 0
	}
	return rand.Intn(len(tickets))
}

func generateTickets(numberOfTickets int) bool {
	fmt.Println("Generating lottery tickets...")
	for i := 0; i < numberOfTickets; i++ {
		ticket := Tickets{Number: randomSixDigitString(i), Name: "", Phone: ""}
		tickets = append(tickets, &ticket)
	}

	if( len(tickets)>0 ){
		return true
	}
	return false
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

func buyTickets() bool {
	var buyTickets int
	var name string
	var phone string
	fmt.Print("\nEnter the name of the ticket buyer:")
	fmt.Scan(&name)
	if !validateName(name) {
		return false
	}	
	fmt.Print("\nEnter the phone number of the ticket buyer:")
	fmt.Scan(&phone)
	if !validatePhone(phone) {
		return false
	}
	fmt.Print("\nHow many tickets would you like to buy?")
	fmt.Scan(&buyTickets)
	if !validNumber(buyTickets) {
		return false
	}

	if( getTicket(buyTickets,name,phone) != false ) {
		fmt.Printf("You have successfully bought ticket with name: %s and phone: %s\n", name, phone)
	}else{
		fmt.Println("Sorry, not enough tickets available.")
	}
	return true
}
