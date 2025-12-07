package main

import (
	"fmt"
)

func calculator(){
	var x,y int
	operator:=""
	print("Enter first number: ")
	fmt.Scan(&x)
	
	print("Enter operator(+,-,* or /): ")
	fmt.Scan(&operator)

	print("Enter second number: ")
	fmt.Scan(&y)

	switch operator {
		case "+":
			fmt.Printf("Sum: %d\n", x+y)
		case "-":
			fmt.Printf("Difference: %d\n", x-y)
		case "*":
			fmt.Printf("Product: %d\n", x*y)
		case "/":
			if y != 0 {
				fmt.Printf("Quotient: %d\n", x/y)
			} else {
				fmt.Println("Error: Division by zero")
			}
	}

}