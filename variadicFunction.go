package main

import "fmt"

func variadicFunction(){
	fmt.Println("Enter how many numbers you want to sum:")
	var n int
	fmt.Scan(&n)

	println(sumOfNumbers(1, 2, 3, 4, 5)) // Output: 15
}

func sumOfNumbers(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}