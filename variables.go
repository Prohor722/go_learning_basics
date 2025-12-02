package main

import "fmt"

func showVariables() {
	a:= 10
	b:= 20
	d:= 30.5
	e:= 19.5333
	// _ = d // to avoid unused variable error
	str:= "Hello, Go!"
	sum := a + b
	sum2 := d + e
	println("Sum is:", sum)
	fmt.Println("Sum of float is:", sum2)	// using fmt package to print float sum
	println("String is:", str)
}
