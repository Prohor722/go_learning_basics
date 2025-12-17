package main

import "fmt"

func getGradeNumber() int {
	var a int
	print("Enter Your Number: ")
	fmt.Scan(&a)

	if(a < 0 || a > 100){
		println("Please enter a valid number between 0 to 100")
		return getGradeNumber()
	}

	return a
}

func gradeCal(){
	var number int = getGradeNumber();

	switch {
	case number >= 90 && number <= 100:
		println("Grade: A+")
	case number >= 80 && number < 90:
		println("Grade: A")
	case number >= 70 && number < 80:
		println("Grade: B")
	case number >= 60 && number < 70:
		println("Grade: C")
	case number >= 50 && number < 60:
		println("Grade: D")
	case number >= 0 && number < 50:
		println("Grade: F")
	default:
		println("Invalid Number!")
	}


}