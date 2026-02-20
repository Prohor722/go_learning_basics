package main

import "fmt"

type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)

func getStatus(status OrderStatus) string{
	switch status{
	case 0:
		return "Received"
	case 1:
		return "Confirmed"
	case 2:
		return "Prepared"
	case 3:
		return "Delivered"
	default:
		return "ERROR!!"
	}
}

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to: ",getStatus(status))
}

func enumExamples() {
	changeOrderStatus(Received)
}