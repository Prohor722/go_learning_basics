package main

import "fmt"

func sequenceFormula() int {
	var n int
	println("Enter the value of n:")
	fmt.Scanln(&n)

	result := n * (n + 1) / 2
	println("The sum of the first", n, "natural numbers is:", result)
}