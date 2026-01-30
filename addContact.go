package main

import (
	"fmt"
	"strings"
)

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