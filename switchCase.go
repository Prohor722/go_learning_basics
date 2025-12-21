package main

import "fmt"

func switchCase() {
	var day int
	print("Enter day number (1-7): ")
	fmt.Scan(&day)
	switch day {
	case 1:
		println("Monday")
	case 2:
		println("Tuesday")
	case 3:
		println("Wednesday")
	case 4:
		println("Thursday")
	case 5:
		println("Friday")
	case 6:
		println("Saturday")
	case 7:
		println("Sunday")
	default:
		println("Invalid day number!")
	}
}

func switchCaseExample2() {
	var mark int
	print("Enter your marks (1-100): ")
	fmt.Scan(&mark)

	switch {
		case mark > 80 && mark <= 100:
			println("Grade A")
		case mark > 60 && mark <= 80:
			println("Grade B")
		case mark > 40 && mark <= 60:
			println("Grade C")
		case mark < 0 || mark > 100:
			println("Invalid marks!")
		default:
			println("Fail")
	}

}

func switchCaseExample3() {
	var cls int
	print("Enter your class: (1-12) ")
	fmt.Scan(&cls)

	switch cls{
	case 1,2,3,4,5:
		println("You are in Primary School")
	case 6,7,8:
		println("You are in Middle School")
	case 9,10:
		println("You are in High School")
	case 11,12:
		println("You are in Senior Secondary School")
	default:
		println("Invalid class!")
	}
}

func switchCaseExample4() {
	var month int
	print("Enter month number (1-12): ")
	fmt.Scan(&month)
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		println("31 days")
	case 4, 6, 9, 11:
		println("30 days")