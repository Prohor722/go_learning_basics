package main

import "fmt"

func addContact() {
	var name string
	var phone string
	var email string

	fmt.Print("Enter contact name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter contact phone: ")
	fmt.Scanln(&phone)

	fmt.Print("Enter contact email: ")
	fmt.Scanln(&email)

	fmt.Printf("Contact added: %s, Phone: %s, Email: %s\n", name, phone, email)
}

func validation(inputType string, value interface{}) bool {
	switch inputType {
	case "name":