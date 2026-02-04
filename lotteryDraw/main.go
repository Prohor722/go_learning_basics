package main

type Tickets struct {
	Number string
	Name string
	Phone string
}
var tickets []*Tickets
var lastSoldTicketIndex int = -1
var numberOfTickets int



func main(){
	drawLottery()
}