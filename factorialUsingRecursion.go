package main

func factorialUsingRecursion(){
	// function to calculate factorial using recursion
	var factorial func(n int) int
	factorial = func(n int) int {
		if n == 0 {
			return 1
		}
		
}