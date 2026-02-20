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


//Another way

type OrderStatusList struct{
	Received string
	Confirmed string
	Prepared string
	Delivered string
}
var status = OrderStatusList{
	Received : "Received",
	Confirmed : "Confirmed",
	Prepared : "Prepared",
	Delivered : "Delivered",
}
func enumExamples() {
	// changeOrderStatus(status.Confirmed)
	changeOrderStatus(Received)
}