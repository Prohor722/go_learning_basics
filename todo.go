package main

import "fmt"

func todo(){
	list := []string{
		"Build a Todo List Application",
		"Implement User Authentication",
		"Add Task Prioritization",
	}
	fmt.Println("Welcome to out Todo List Application!!")
	i := 1
	for _, item := range list {
		fmt.Printf("%d. %s\n", i, item)
		i++
	}

}