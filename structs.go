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
}

func strucksExamples(){
	order := Order{
		id: "1",
		ammount: 55.10,
		status: "received",
		createdAt: time.Now(),
	}

	fmt.Println("order struct: ",order)

}