package main

import (
	"fmt"
	"math"
)

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
	sum := a * (int(math.Pow(float64(r), float64(n))) - 1) / (r - 1)
	println("The sum of the geometric progression is:", sum)
}

func fibonacciSeries(){
	var n int
	print("Enter the number of terms:")
	fmt.Scanln(&n)
	a, b := 0, 1
	println("Fibonacci Series:")
	for i := 0; i < n; i++ {
		print(a, " ")
		a, b = b, a+b
	}
}

func sumOfFirstNNumbers(){
	var n int
	print("Enter a positive integer:")
	fmt.Scanln(&n)
	sum := n * (n + 1) / 2
	println("Sum of the first", n, "natural numbers is:", sum)
}

func SumOfSquares(){
	var n int
	print("Enter a positive integer:")
	fmt.Scanln(&n)
	sum := n * (n + 1) * (2*n + 1) / 6
	println("Sum of squares of the first", n, "natural numbers is:", sum)
}

func sumOfCubes(){
	var n int
	print("Enter a positive integer:")
	fmt.Scanln(&n)
	sum := (n * (n + 1) / 2) * (n * (n + 1) / 2)
	println("Sum of cubes of the first", n, "natural numbers is:", sum)
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func permutation(){
	var n, r int
	print("Enter n (total items):")
	fmt.Scanln(&n)
	print("Enter r (items to arrange):")
	fmt.Scanln(&r)
	numerator := factorial(n)
	denominator := factorial(n - r)
	result := numerator / denominator
	println("Number of permutations (P(n,r)) is:", result)
}

func combination(){
	var n, r int
	print("Enter n (total items):")
	fmt.Scanln(&n)
