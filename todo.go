package main

import "fmt"
var list = []string{
		"Build a Todo List Application",
		"Implement User Authentication",
		"Add Task Prioritization",
	}

func readAllTodo(){
	
	i := 1
	for _, item := range list {
		fmt.Printf("%d. %s\n", i, item)
		i++
	}
}

func addTask(task string){
	list = append(list, task)
	fmt.Println("Task added:", task)
}

func todo(){
	fmt.Println("Welcome to out Todo List Application!!")
	readAllTodo()

}