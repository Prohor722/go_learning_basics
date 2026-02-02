package main

import "fmt"

func drawLottery(winningNumbers []int, userNumbers []int, contact *Contact) {
	matchingNumbers := []int{}
	for _, userNum := range userNumbers {
		for _, winNum := range winningNumbers {
			if userNum == winNum {
				matchingNumbers = append(matchingNumbers, userNum)
			}
		}
	}
	if len(matchingNumbers) == 0 {
		fmt.Println("No matching numbers. Better luck next time!")
	}
	if len(matchingNumbers) > 0 {
		fmt.Printf("Congratulations! You have %d matching numbers: %v\n", len(matchingNumbers), matchingNumbers)
	}
	var name, phone, email string
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
	fmt.Printf("Contact updated: %s, Phone: %s, Email: %s\n", name, phone, email)
	fmt.Println("Current Contacts:")
	printContacts()
}

func addContact() {