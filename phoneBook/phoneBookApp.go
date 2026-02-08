package main

import (
	"fmt"
)

func phoneBookApp() {
	for {
		var choice int
		printMenuOptions()
		
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
			index = getIndex()
			editContact(contacts[index-1], index-1)
		
		case 4:
			fmt.Println("Which contact would you like to delete?")
			printContacts()
			var index int
			index = getIndex()
			deleteContact(index - 1)
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
