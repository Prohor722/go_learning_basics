package main

import "fmt"

func changeNum(num *int) {
	*num = 5
	fmt.Println("in changeFunc: ",num)
}

func pointersExample() {
	num := 1

	changeNum(&num)

	fmt.Println("After changeNum in main: ", num)
}