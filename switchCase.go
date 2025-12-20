package main

func switchCase(){
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