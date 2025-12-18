package main

import "fmt"

func bmiCal() {
	// BMI = weight (kg) / height^2 (m^2)
	var weight, height, bmi float64;
	print("Enter your weight in kg: ")
	fmt.Scan(&weight)
	println("In which parameter you want your height to be calculated?")
	println("1.Feet/Inches \n2.Meters (Select 1 or 2): ")
	var choice int
	fmt.Scan(&choice)

	if choice<1 || choice>2 {
		println("Invalid choice !! Please select again.")
		bmiCal()
	}

	if choice == 1 {
		var feet, inches float64
		print("Enter your height - Feet: ")
		fmt.Scan(&feet)
		print("Inches: ")
		fmt.Scan(&inches)
		height = (feet * 12 + inches) * 0.0254 // converting to meters
	} else {
		print("Enter your height in meters: ")
		fmt.Scan(&height)
	}

	bmi = weight / (height * height)

	println("Your BMI is:", bmi)
	if bmi < 18.5 {
		println("You are underweight.")
	} else if bmi >= 18.5 && bmi < 24.9 {
		println("You have a normal weight.")
	} else if bmi >= 25 && bmi < 29.9 {
		println("You are overweight.")
	}else {
		println("You are obese.")
	}
	
}