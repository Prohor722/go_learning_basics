package main

import (
	"fmt"
)

func formatVariables() {
	var decimalNumber int

	print("Enter a decimal number: ")
	fmt.Scan(&decimalNumber)

	fmt.Printf("\n You number in decimal: %d", decimalNumber) // decimal format
	fmt.Printf("\n Your number in binary: %b", decimalNumber) // binary format
	fmt.Printf("\n Your number in hexadecimal: %x\n", decimalNumber) // hexadecimal format
	fmt.Printf("Your number in octal: %o\n", decimalNumber) // octal format
	fmt.Printf("Your number with leading zeros: %08d\n", decimalNumber) // leading zeros
	fmt.Printf("Your number with sign: %+d\n", decimalNumber) // with sign
	fmt.Printf("Your number in width 10: %10d\n", decimalNumber) // width 10
	fmt.Printf("Your number left aligned in width 10: %-10d\n", decimalNumber) // left aligned width 10
	
}