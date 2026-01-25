package main

import "fmt"
var list = map[string]string{
		"App":"Build a Todo List Application",
		"Authentication":"Implement User Authentication",
		"Prioritization":"Add Task Prioritization",
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
		fmt.Println("Invalid task index.")
		return
	}
	i := 0
	for key := range list {
		if i == index {
			delete(list, key)
			fmt.Println("Task removed.")
			return
		}
		i++
	}
}

func scanTask() (string,string) {
	var title string
	fmt.Print("Enter the task title: ")
	fmt.Scanln(&title)
	fmt.Print("Enter Task Details: ")
	var details string
	fmt.Scanln(&details)
	return title, details
}

func addTask(title string, details string){
	list[title] = details
	fmt.Println("Task added.")
}

func editTask(index int, newDetails string){
	if index < 0 || index >= len(list) {
		fmt.Println("Invalid task index.")
		return
	}
	i := 0
	for key := range list {
		if i == index {
			list[key] = newDetails
			fmt.Println("Task updated.")
			return
		}
		i++
	}
}

func editTaskTitle(index int, newTitle string){
	if index < 0 || index >= len(list) {
		fmt.Println("Invalid task index.")
		return
	}
	i := 0
	for key, value := range list {
		if i == index {
			delete(list, key)
			list[newTitle] = value
			fmt.Println("Task title updated.")
			return
		}
		i++
	}
}

func todo(){
	fmt.Println("Welcome to out Todo List Application!!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. View all tasks")
	fmt.Println("2. Add a new task")
	fmt.Println("3. Remove a task")

}