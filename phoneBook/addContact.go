package main

import (
	"fmt"
)

func addContact() []*Contact {
	var name string
	var phone string
	var email string

	name = getName()
	phone = getPhone()
	email = getEmail()

	if !contactExists(name, phone) {
		contacts = append(contacts, &Contact{Name: name, Phone: phone, Email: email})
		fmt.Printf("Contact added! Name: %s, Phone: %s, Email: %s\n", name, phone, email)
		fmt.Println("Current Contacts:")
		printContacts()
	}
	return contacts
}
