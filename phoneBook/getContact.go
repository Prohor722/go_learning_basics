package main

import (
	"fmt"
)

func getName() string {
	var name string
	fmt.Print("Enter contact name: ")
	fmt.Scanln(&name)
	if(!validateName(name)) {
		fmt.Println("Invalid name. Please try again.")
		return getName()
	}
	return name
}

func getPhone() string {
	var phone string
	fmt.Print("Enter contact phone: ")
	fmt.Scanln(&phone)
	if(!validatePhone(phone)) {
		fmt.Println("Invalid phone number. Please try again.")
		return getPhone()
	}
	return phone
}

func getEmail() string {
	var email string
	fmt.Print("Enter contact email: ")
	fmt.Scanln(&email)
	if(!validateEmail(email)) {
		fmt.Println("Invalid email address. Please try again.")
		return getEmail()
	}
	return email
}