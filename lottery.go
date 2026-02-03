package main

import "fmt"

type Tickets struct {
	Numbers []int
}

tickets := []Tickets{}

func drawLottery(winningNumbers []int, userNumbers []int, contact *Contact) {

}

func generateTickets(){
	fmt.Println("Generating lottery tickets...")
	for i := 0; i < 5; i++ {
		ticket := Tickets{Numbers: []int{i, i + 1, i + 2, i + 3, i + 4}}
		tickets = append(tickets, ticket)
	}
}