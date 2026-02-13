package main

import "fmt"

func variadicFunction(){
	fmt.Println("Enter how many numbers you want to sum:")
	var n int
	fmt.Scan(&n)
	fmt.Println("Enter the numbers:")
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	println(sumOfNumbers(1, 2, 3, 4, 5)) // Output: 15
}

func sumOfNumbers(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}