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
