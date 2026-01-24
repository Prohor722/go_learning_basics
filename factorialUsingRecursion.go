package main

import "fmt"

func factorialUsingRecursion() {
	// function to calculate factorial using recursion
	var factorial func(n int) int
	factorial = func(n int) int {
		if n == 0 {
			return 1
		}
		return n * factorial(n-1)
	}
	var num int
	print("Please enter the number:")
	fmt.Scan(&num)
	result := factorial(num)
	println("\nFactorial of", num, "is", result)
}