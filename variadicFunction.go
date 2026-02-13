package main

import "fmt"

func variadicFunction(){
	fmt.Println("Enter your first name:")
	var firstName string
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	var lastName string
	fmt.Scan(&lastName)

	printFullName(firstName, lastName)

	fmt.Println("Enter how many numbers you want to sum:")
	var n int
	fmt.Scan(&n)
	fmt.Println("Enter the numbers:")
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	println(sumOfNumbers(nums...)) // Output: sum of all numbers in nums slice
}

func sumOfNumbers(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func printFullName(firstName, lastName string) {
	fmt.Printf("Full Name: %s %s\n", firstName, lastName)
}