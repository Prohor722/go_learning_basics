package main

import (
	"fmt"
	"strings"
)

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
	deletedContacts = append(deletedContacts, contacts[index])
	contacts = append(contacts[:index], contacts[index+1:]...)
	fmt.Println("Contact deleted successfully.")
}