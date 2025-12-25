package main

import "fmt"

func sequenceFormula() {
	var n int
	print("Enter the value of n:")
	fmt.Scanln(&n)

	result := n * (n + 1) / 2
	println("The sum of the first", n, "natural numbers is:", result)
}

func arithmeticProgression(a, d, n int) int {
	return n * (2*a + (n-1)*d) / 2
}