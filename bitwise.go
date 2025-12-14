package main

import "fmt"

func bitwise(){
	a := 60        // 60 = 0011 1100
	b := 13 	  // 13 = 0000 1101

	fmt.Printf("a & b = %d\n", a & b)   // 12 = 0000 1100
	fmt.Printf("a | b = %d\n", a | b)   // 61 = 0011 1101
}