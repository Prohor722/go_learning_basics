package main

import "fmt"


func printSlice[T int | string](items [] T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func genericExample() {
	nums := []int{1,2,3,4,5}
	names := []string{"golang","ts","js"}
	printSlice(names)
	printSlice(nums)
}