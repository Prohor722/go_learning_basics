package main

import (
	"fmt"
	"strings"
)

func validateChoice(choice int) bool {
	if choice < 1 || choice > 6 {
		fmt.Println("Invalid choice. Please try again.")
		return false
	}
	return true
}

func validateName(name string) bool {
	if len(name) < 2 {
		fmt.Println("Invalid name. Please try again.")
		return false
	}
	return true
}

func validatePhone(phone string) bool {
	if len(phone) < 7 {
		fmt.Println("Invalid phone number. Please try again.")
		return false
	}
	return true
}

func validateEmail(email string) bool {
	if len(email) < 5 || !strings.Contains(email, "@") {
		fmt.Println("Invalid email address. Please try again.")
		return false
	}
	return true
}

func contactExists(name string, phone string) bool {
	for _, contact := range contacts {
		if contact.Name == name && contact.Phone == phone {
			fmt.Printf("Contact already exists! Name: %s, Phone: %s, Email: %s\n", contact.Name, contact.Phone, contact.Email)
			return true
		}
	}
	return false
}

func validateIndex(index int) bool {
	if index < 1 || index > len(contacts) {
		return false
	}
	return true
}