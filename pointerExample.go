package main

import "fmt"

type Person struct {
	name string
	age  int
	address string
}

func pointerExample() {
	person1 := Person{name: "Alice", age: 30, address: "123 Main St"}
	fmt.Printf("Before modification: %+v\n", person1)
}