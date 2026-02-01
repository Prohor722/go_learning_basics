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
var contacts []*Contact

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
	contacts = append(contacts, &Contact{Name: name, Phone: phone, Email: email})
	fmt.Printf("Contact added: %s, Phone: %s, Email: %s\n", name, phone, email)
	fmt.Println("Current Contacts:")
	// fmt.Println(contacts)
	printContacts()
}

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

func deleteContact(index int) {
	fmt.Printf("Deleting contact: %s\n", contacts[index].Name)
	fmt.Println("Are you sure?")
	var confirmation string
	fmt.Print("Type 'yes' to confirm: ")
	fmt.Scanln(&confirmation)
	if strings.ToLower(confirmation) != "yes" {
		fmt.Println("Deletion cancelled.")
		return
	}
	contacts = append(contacts[:index], contacts[index+1:]...)
	fmt.Println("Contact deleted successfully.")
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

func contactExists(name string) bool {
	for _, contact := range contacts {
		if contact.Name == name {
			return true
		}
	}
	return false
}

func addContactIfNotExists(name string, phone string, email string) []*Contact {
	if !contactExists(name) {
		contacts = append(contacts, &Contact{Name: name})
		fmt.Printf("Contact %s added successfully.\n", name)
	}
	return contacts
}

func printContacts() {
	fmt.Println("Contact List:")
	for i, contact := range contacts {
		fmt.Printf("%d. %s\n", i+1, contact.Name)
		fmt.Printf("   Phone: %s\n", contact.Phone)
		fmt.Printf("   Email: %s\n", contact.Email)
	}
}

func phoneBookApp() {
	for {
		var choice int
		fmt.Println("\nPhone Book Menu:")
		fmt.Println("1. Add Contact")
		fmt.Println("2. View Contacts")
		fmt.Println("3. Update Contact")
		fmt.Println("4. Delete contact")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			addContact()
		case 2:
			printContacts()
		case 3:
			printContacts()
			fmt.Println("Which contact would you like to edit?")
			var index int
			// fmt.Print("Enter contact number: ")
			fmt.Scanln(&index)
			if index > 0 && index <= len(contacts) {
				// In a real app, we would retrieve the actual Contact struct by index
				contact := contacts[index-1]
				editContact(contact, index-1)
			} else {
				fmt.Println("Invalid contact number.")
			}

		case 4:
			fmt.Println("Which contact would you like to delete?")
			printContacts()
			var index int
			fmt.Print("\nEnter contact serial number: ")
			fmt.Scanln(&index)
			if index <= 0 || index > len(contacts) {
				fmt.Println("Invalid contact number.")
				continue
			}
			deleteContact(index-1)
		case 5:
			fmt.Println("Exiting Phone Book App.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}