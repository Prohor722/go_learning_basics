package main

import (
	"fmt"
	"strings"
)

type Contact struct {
	Name  string
	Phone string
	Email string
}

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
	fmt.Printf("Contact added: %s, Phone: %s, Email: %s\n", name, phone, email)
}

func editContact(contacts []*Contact, contact *Contact, index int) {
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



func contactDetailsValidation(inputType string, value interface{}) bool {
	switch inputType {
	case "name":
		name, ok := value.(string)
		if !ok || len(name) < 2 {
			return false
		}
		return true
	case "phone":
		phone, ok := value.(string)
		if !ok || len(phone) < 7 {
			return false
		}
		return true
	case "email":
		email, ok := value.(string)
		if !ok || len(email) < 5 || !strings.Contains(email, "@") {
			return false
		}
		return true
	default:
		return false
	}
}

func contactExists(contacts []*Contact, name string) bool {
	for _, contact := range contacts {
		if contact.Name == name {
			return true
		}
	}
	return false
}

func addContactIfNotExists(contacts []*Contact, name string) []*Contact {
	if !contactExists(contacts, name) {
		contacts = append(contacts, &Contact{Name: name})
		fmt.Printf("Contact %s added successfully.\n", name)
	}
	return contacts
}

func printContacts(contacts []*Contact) {
	fmt.Println("Contact List:")
	for i, contact := range contacts {
		fmt.Printf("%d. %s\n", i+1, contact.Name)
	}
}

func phoneBookApp() {
	contacts := []*Contact{}
	for {
		var choice int
		fmt.Println("\nPhone Book Menu:")
		fmt.Println("1. Add Contact")
		fmt.Println("2. View Contacts")
		fmt.Println("3. Edit Contact")
		fmt.Println("4. Delete contact")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			addContact()
		case 2:
			printContacts(contacts)
		case 3:
			fmt.Println("Which contact would you like to edit?")
			printContacts(contacts)
			var index int
			fmt.Print("Enter contact number: ")
			fmt.Scanln(&index)
			if index > 0 && index <= len(contacts) {
				// In a real app, we would retrieve the actual Contact struct by index
				contact := contacts[index-1]
				editContact(contacts, contact, index-1)
			} else {
				fmt.Println("Invalid contact number.")
			}

		case 4:
			// deleteContact()
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}