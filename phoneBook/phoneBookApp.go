package main

import (
	"fmt"
)

func phoneBookApp() {
	for {
		var choice int
		fmt.Println("\nPhone Book Menu:")
		fmt.Println("1. Add Contact")
		fmt.Println("2. View Contacts")
		fmt.Println("3. Update Contact")
		fmt.Println("4. Delete contact")
		fmt.Println("5. View Deleted Contacts")
		fmt.Println("6. Exit")
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
			fmt.Scanln(&index)
			
			if index > 0 && index <= len(contacts) {
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
			printDeletedContacts()
		case 6:
			fmt.Println("Exiting Phone Book App.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}