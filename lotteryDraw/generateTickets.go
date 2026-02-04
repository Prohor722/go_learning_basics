package main

import (
	"fmt"
	"math/rand"
	"time"
)

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