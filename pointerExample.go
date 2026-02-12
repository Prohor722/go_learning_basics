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

func printPersonInfo(p Person) {
	fmt.Printf("Name: %s, Age: %d, Address: %s\n", p.name, p.age, p.address)
}

func pointerExample() {
	person1 := Person{name: "Alice", age: 30, address: "123 Main St"}
	printPersonInfo(person1)
	person1.incresseAge(5)
	fmt.Println("After modification:")
	printPersonInfo(person1)
	person2 := person1.namehonorific("Dr.")
	fmt.Println("After adding honorific:")
	printPersonInfo(person2)
}