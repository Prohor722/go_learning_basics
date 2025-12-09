package main

import (
	"fmt"
)

func calculator(){
	var x,y int
	operator:=""
	print("Enter first number: ")
	fmt.Scan(&x)
	
	print("Enter operator(+,-,* or /): ")
	fmt.Scan(&operator)

	print("Enter second number: ")
	fmt.Scan(&y)

	switch operator {
		case "+":
			fmt.Printf("Sum: %d\n", x+y)
		case "-":
			fmt.Printf("Difference: %d\n", x-y)
		case "*":
			fmt.Printf("Product: %d\n", x*y)
		case "/":
			if y != 0 {
				fmt.Printf("Quotient: %d\n", x/y)
			} else {
				fmt.Println("Error: Division by zero")
			}
	}

}

func areaCalculator() {
	var radius float64

	print("Calculate area of rectangle, tringle or circle? (r/t/c): ")
	var choice string
	fmt.Scan(&choice)

	if choice == "r" {	
		print("Enter the radius of the circle: ")
		fmt.Scan(&radius)
		area := areaOfCircle(radius)
		fmt.Printf("The area of the circle with radius %.2f is %.2f\n", radius, area)
	} else if choice == "c" {
		var length, width float64
		print("Enter the length of the rectangle: ")
		fmt.Scan(&length)
		print("Enter the width of the rectangle: ")
		fmt.Scan(&width)
		area := areaOfRectangle(length, width)
		fmt.Printf("The area of the rectangle with length %.2f and width %.2f is %.2f\n", length, width, area)
	} else if choice == "t" {
		var base, height float64
		print("Enter the base of the triangle: ")
		fmt.Scan(&base)
		print("Enter the height of the triangle: ")
		fmt.Scan(&height)
		area := areaOfTriangle(base, height)
		fmt.Printf("The area of the triangle with base %.2f and height %.2f is %.2f\n", base, height, area)
	} else {
		fmt.Println("Invalid choice. Please enter 'r' for rectangle or 'c' for circle.")
	}
}

func areaOfCircle(radius float64) float64 {
	const pi = 3.14159
	return pi * radius * radius
}

func areaOfRectangle(length, width float64) float64 {
	return length * width
}

func areaOfTriangle(base, height float64) float64 {
	return 0.5 * base * height
}