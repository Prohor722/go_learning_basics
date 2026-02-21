package main

import "fmt"


func printSlice[T int | string](items [] T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

type stack[T int|string] struct{
	elements [] T
}

func genericExample() {
	myStack := stack[string]{
		elements: []string{
			"Hello",
			"Brother",
		},
	}

	// myStack.elements = []int{	//only can reassign the same type data
	// 	1,
	// 	2,
	// }

	fmt.Println(myStack)

	nums := []int{1,2,3,4,5}
	names := []string{"golang","ts","js"}
	printSlice(names)
	printSlice(nums)
}