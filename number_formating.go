package main

import (
	"fmt"
)

func formatVariables() {
	var decimalNumber int

	print("Enter a decimal number: ")
	fmt.Scan(&decimalNumber)

	fmt.Printf("\n You have entered: %b", decimalNumber) // binary format
}