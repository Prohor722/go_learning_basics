package main

import "fmt"

type Person struct {
	name string
	age  int
	address string
}

func (p *Person) incresseAge (incresseBy int) {
	p.age += incresseBy
}

func (p Person) namehonorific(honorific string) Person {
	p.name = honorific + " " + p.name
	return p
}

func pointerExample() {
	person1 := Person{name: "Alice", age: 30, address: "123 Main St"}
	fmt.Printf("Before modification: %+v\n", person1)
	person1.incresseAge(5)
	fmt.Printf("After modification: %+v\n", person1)
	person2 := person1.namehonorific("Dr.")
	fmt.Printf("After adding honorific: %+v\n", person2)
}