package main

import "fmt"

func changeNum(num int) {
	num = 5
	fmt.Println("in change num: ",num)
}

func pointersExample() {
	num := 1

	changeNum(num)

	fmt.Println("After change num in main: ", num)
}