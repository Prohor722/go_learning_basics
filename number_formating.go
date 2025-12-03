package main

import (
	"fmt"
)

func formatVariables() {
	var decimalNumber int

	print("Enter a decimal number: ")
	fmt.Scan(&decimalNumber)

	println("\n You have entered: ", decimalNumber)
}