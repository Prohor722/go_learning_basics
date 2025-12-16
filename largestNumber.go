package main

import "fmt"

func getNumber(serial string) int {
	var a int
	fmt.Printf("\nEnter %v Number: ", serial)
	fmt.Scan(&a)
	return a
}
func largestNumber(){
	num1 := getNumber("1st")
	num2 := getNumber("2nd")

	if(num1 > num2){
		println("\nLargest Number is: ", num1)
	}else if(num1 == num2){
		fmt.Printf("\nBoth Numbers %d and %d are Equal", num1, num2)
	}else{
		println("\nLargest Number is: ", num2)
	}
}