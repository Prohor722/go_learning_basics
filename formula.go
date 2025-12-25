package main

import "fmt"

func sequenceFormula() {
	var n int
	print("Enter the value of n:")
	fmt.Scanln(&n)

	result := n * (n + 1) / 2
	println("The sum of the first", n, "natural numbers is:", result)
}

func arithmeticProgression() {
	var a, d, n int
	print("Enter the first term (a):")
	fmt.Scanln(&a)
	print("Enter the common difference (d):")
	fmt.Scanln(&d)
	print("Enter the number of terms (n):")
	fmt.Scanln(&n)

	lastTerm := a + (n-1)*d
	sum := n * (a + lastTerm) / 2
	println("The sum of the arithmetic progression is:", sum)
}

func GeometricProgression(){
	var a, r, n int
	print("Enter the first term (a):")
	fmt.Scanln(&a)
	print("Enter the common ratio (r):")
	fmt.Scanln(&r)
	print("Enter the number of terms (n):")
	fmt.Scanln(&n)

	sum := a * (pow(r, n) - 1) / (r - 1)

}