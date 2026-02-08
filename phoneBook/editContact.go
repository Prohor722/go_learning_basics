package main

import (
	"fmt"
)

func editContact(contact *Contact, index int) {
	var name string
	var phone string
	var email string

	

	contact.Name = name
	contact.Phone = phone
	contact.Email = email

	contacts[index] = contact
	
	fmt.Printf("Contact updated: %s, Phone: %s, Email: %s\n", contact.Name, contact.Phone, contact.Email)
}
