package main

import (
	"fmt"
)

type Tickets struct {
	Number string
	Name string
	Phone string
}
var tickets []*Tickets
var lastSoldTicketIndex int = -1
var totalLotteryTickets int

func main(){
	startGame()
}
