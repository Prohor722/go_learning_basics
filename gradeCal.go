package main

import "fmt"

func getGradeNumber() int {
	var a int
	print("Enter Number: ")
	fmt.Scan(&a)
	return a
}

func gradeCal(){
	var number int = getGradeNumber();

	switch {
	case number >= 90 && number <= 100:
		println("Grade: A+")
	case number >= 80 && number < 90:
		println("Grade: A")


}