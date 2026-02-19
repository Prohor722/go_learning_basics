package main

import (
	"fmt"
	"time"
)

//order struct
type Order struct{
	id string
	ammount float32
	status string
	createdAt time.Time // nanosecond precision
	updatedAt time.Time
}

func (o Order) changeStatus(status string){
	o.status = status
}

func strucksExamples(){
	order := Order{
		id: "1",
		ammount: 55.10,
		status: "received",
		createdAt: time.Now(),
	}

	var orders []Order

	orders = append(orders, order)

	order.updatedAt = time.Now()

	fmt.Println("order struct: ",order)
	fmt.Println("orders array: ",orders)


	order.changeStatus("on processing")

	fmt.Println("After changing status: ", order)

}