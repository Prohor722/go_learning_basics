package main

imort "fmt"

func addContact() {
	var name string
	var phone string
	var email string

	fmt.Print("Enter contact name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter contact phone: ")
	fmt.Scanln(&phone)
}