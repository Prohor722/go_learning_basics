package main

import (
	"fmt"
)

func editContact(contact *Contact, index int) {
	var name string
	var phone string
	var email string

	fmt.Print("Enter new contact name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter new contact phone: ")
	fmt.Scanln(&phone)
	fmt.Print("Enter new contact email: ")
	fmt.Scanln(&email)

	contactDetailsValidation("name", name)
	contactDetailsValidation("phone", phone)
	contactDetailsValidation("email", email)

	if(!contactDetailsValidation("name", name)) {
		fmt.Println("Invalid name. Please try again.")
		return
	}
	if(!contactDetailsValidation("phone", phone)) {
		fmt.Println("Invalid phone number. Please try again.")
		return
	}
	if(!contactDetailsValidation("email", email)) {
		fmt.Println("Invalid email address. Please try again.")
		return
	}

	contact.Name = name
	contact.Phone = phone
	contact.Email = email

	contacts[index] = contact
	
	fmt.Printf("Contact updated: %s, Phone: %s, Email: %s\n", contact.Name, contact.Phone, contact.Email)
}
