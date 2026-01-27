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

func variableActAsConstant() {
	fmt.Println("Value of pi is:", pi())
}