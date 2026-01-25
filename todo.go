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

func removeTask(index int){
	if index < 0 || index >= len(list) {
		fmt.Println("Invalid index")
		return
	}
	removedTask := list[index]
	list = append(list[:index], list[index+1:]...)
	fmt.Println("Task removed:", removedTask)
}

func addTask(task string){
	list = append(list, task)
	fmt.Println("Task added:", task)
}

func todo(){
	fmt.Println("Welcome to out Todo List Application!!")
	readAllTodo()
	fmt.Println("\nAdding a new task: 'Learn Go Concurrency'")
	addTask("Learn Go Concurrency")
	fmt.Println("\nUpdated Todo List:")
	readAllTodo()
	fmt.Println("\nRemoving the second task...")
	removeTask(1)
	fmt.Println("\nUpdated Todo List:")
	readAllTodo()

}