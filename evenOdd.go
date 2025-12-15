package main

import "fmt"

func evenOdd(){
	var num int
	print("Enter an integer: ")
	fmt.Scanf("%d", &num)

	if num%2 == 0 {
		println(num, "is even")
	} else {
		println(num, "is odd")
	}
}