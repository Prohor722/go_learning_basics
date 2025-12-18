package main

func bmiCal() {
	// BMI = weight (kg) / height^2 (m^2)
	var weight, height, bmi float64;
	print("Enter your weight in kg: ")
	fmt.Scan(&weight)
	print("Enter your height in meters: ")
	fmt.Scan(&height)

	bmi = weight / (height * height)

	println("Your BMI is:", bmi)
	
}