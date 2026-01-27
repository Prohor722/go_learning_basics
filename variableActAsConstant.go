package main

import "fmt"

func pi() float64 {
	return 3.1416;
}

func name() string {
	return "Prohor";
}

func age() int {
	return 250;
}

func height() float64 {
	return 9.9;
}

func variableActAsConstant() {
	fmt.Println("Value of pi is:", pi())
	fmt.Println("Name is:", name())
	fmt.Println("Age is:", age())
	fmt.Println("Height is:", height())
}