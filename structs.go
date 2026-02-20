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

func (o *Order) changeStatus(status string){
	o.status = status
}

func (o Order) getAmmount() float32{
	return o.ammount
}

// Constructor
func NewOrder(id string, ammount float32, status string) *Order{
	myOrder := Order{
		id: id,
		ammount: ammount,
		status: status,
		createdAt: time.Now(),
	}

	return &myOrder
}

func strucksExamples(){
	// order := Order{
	// 	id: "1",
	// 	ammount: 55.10,
	// 	status: "received",
	// 	createdAt: time.Now(),
	// }
	order := NewOrder("1",55.10,"received")

	// var orders []Order

	// orders = append(orders, order)

	order.updatedAt = time.Now()

	fmt.Println("order struct: ",order)
	// fmt.Println("orders array: ",orders)


	order.changeStatus("on-processing")

	fmt.Println("After changing status: ", order)

	fmt.Println("Printing ammount: ",order.getAmmount())

}


/* STUCT EMBEDDING (like inheritence in oop) */

type Civilian struct{
	name string
	age int
}

type Police struct{
	rank string
	post string
	Civilian
}

func struckEmbedingExample(){
	newPeople := Civilian {
			name:"Billu Mia",
			age: 32,
		}
	newOfficer := Police{
		rank: "ASP",
		post: "Dhaka",
		Civilian: newPeople,
		// Civilian: Civilian {		//Alternative
		// 	name:"Billu Mia",
		// 	age: 32,
		// },
	}

	fmt.Println("New Officer: ",newOfficer)
}