package main

import (
	"fmt"
)

func printContacts() {
	fmt.Println("Contact List:")
	for i, contact := range contacts {
		fmt.Printf("%d. %s\n", i+1, contact.Name)
		fmt.Printf("   Phone: %s\n", contact.Phone)
		fmt.Printf("   Email: %s\n", contact.Email)
	}
}

func printDeletedContacts() {
	fmt.Println("Deleted Contacts:")
	for i, contact := range deletedContacts {
		fmt.Printf("%d. %s\n", i+1, contact.Name)
		fmt.Printf("   Phone: %s\n", contact.Phone)
		fmt.Printf("   Email: %s\n", contact.Email)
	}
}