package main

import "fmt"

func getNumber(serial string) int {
	var a int
	fmt.Printf("\nEnter %s Number: ", serial)
	fmt.Scan(&a)
	return a
}
func largestNumber(){
	num1 := getNumber("1st")
	num2 := getNumber("2nd")

	if(num1 > num2){
		println("\nLargest Number is: ", num1)
	} else{
		println("\nLargest Number is: ", num2)
	}
}