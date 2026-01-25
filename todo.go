package main

import "fmt"

func readAllTodo(){
	list := []string{
		"Build a Todo List Application",
		"Implement User Authentication",
		"Add Task Prioritization",
	}
	i := 1
	for _, item := range list {
		fmt.Printf("%d. %s\n", i, item)
		i++
	}
}

func todo(){
	
	fmt.Println("Welcome to out Todo List Application!!")
	readAllTodo()

}