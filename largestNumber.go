package main

import "fmt"

func getNumber() int {
	var a int
	println("\nEnter Number: ")
	fmt.Scan("%d", &a)
	return a
}
func largestNumber(){
	num1 := getNumber()
	num2 := getNumber()

	if(num1 > num2){
		println("\nLargest Number is: ", num1)
	}
}