package main

import (
	"fmt"
)

func validateName(name string) bool {
	var minLength = 2
	if len(name) < minLength {
		fmt.Println("Name is too short.(Min:", minLength, ")")
		return false
	}
	return true
}
func validatePhone(phone string) bool {
	var minLength = 5
	if len(phone) < minLength {
		fmt.Printf("Phone number is too short.(Min: %d)\n", minLength)
		return false
	}
	return true
}

func validNumber(number int) bool {
	if( number < 1 ){
		fmt.Println("Number must be greater than zero.")
		return false
	}
	return true
}

func getPhone() string {
	var phone string
	fmt.Print("Enter your phone number:")
	fmt.Scan(&phone)
	validatePhone(phone)
	return phone
}

func getName() string {
	var name string
	fmt.Print("Enter your name:")