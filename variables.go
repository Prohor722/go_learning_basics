package main

import (
	"fmt"
)

func showVariables() {
	a:= 10
	var b int = 20
	d:= 30.5
	var e float64 = 19.5333
	var f float32 = 12.33
	var g float32 = 45.66
	const abc = 100
	const pi float64 = 3.14159
	const str1 string = "Constant String"

	var sum3 float32 = f + g
	// _ = d // to avoid unused variable error
	str:= "Hello, Go!"
	var str2 string = "Go programming."
	sum := a + b
	sum2 := d + e
	println("Sum is:", sum)
	fmt.Println("Sum of float is:", sum2)	// using fmt package to print float sum
	fmt.Println("Sum of float32 is:", sum3)
	println("String is:", str)
	fmt.Printf("One line String is: %v\n", str)
	println(str2)

	println("this is a variable: ",a)
	fmt.Printf("this is a variable: %d\n", b)

	fmt.Printf("Constant abc: %d, pi: %.5f, str1: %s\n", abc, pi, str1)
}
