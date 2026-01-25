package main

import (
	"bufio"
	"fmt"
	"os"
)

var list = map[string]string{
	"App":            "Build a Todo List Application",
	"Authentication": "Implement User Authentication",
	"Prioritization": "Add Task Prioritization",
}

func readAllTodo() {
	if len(list) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	fmt.Println("Your Todo List:")
	i := 1
	for key, value := range list {
		fmt.Printf("%d. %s: %s\n", i, key, value)
		i++
	}
	fmt.Println()
}

func removeTask(index int) {
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
	fmt.Println()
}

func scanTask() (string, string) {
	scanner := bufio.NewScanner(os.Stdin)

	var title string
	fmt.Print("Enter the task title: ")
	scanner.Scan()
	title = scanner.Text()

	fmt.Print("Enter Task comment and press enter when complete: ")
	var details string
	scanner.Scan()
	details = scanner.Text()

	return title, details
}

func addTask() {
	title, details := scanTask()
	list[title] = details
	fmt.Println("Task added.")
	fmt.Println()
}

func editTask(index int, newDetails string) {
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
	fmt.Println()
}

func editTaskTitle(index int, newTitle string) {
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
	fmt.Println()
}

func todo() {
	fmt.Println("Welcome to out Todo List Application!!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. View all tasks")
	fmt.Println("2. Add a new task")
	fmt.Println("3. Remove a task")
	fmt.Println("4. Edit a task")
	fmt.Println("5. Edit a task title")
	fmt.Println("0. Exit")

	var choice int
	fmt.Print("Enter your choice (1-5): ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		readAllTodo()
		todo()
	case 2:
		addTask()
		todo()
	case 3:
		var index int
		fmt.Print("Enter the task index to remove: ")
		fmt.Scanln(&index)
		removeTask(index - 1)
		todo()
	case 4:
		var index int
		fmt.Print("Enter the task index to edit: ")
		fmt.Scanln(&index)
		fmt.Print("Enter new task details: ")
		var newDetails string
		fmt.Scanln(&newDetails)
		editTask(index-1, newDetails)
		todo()
	case 5:
		var index int
		fmt.Print("Enter the task index to edit title: ")
		fmt.Scanln(&index)
		fmt.Print("Enter new task title: ")
		var newTitle string
		fmt.Scanln(&newTitle)
		editTaskTitle(index-1, newTitle)
		todo()
	case 0:
		fmt.Println("Exiting Todo List Application. Goodbye!")
		return
	default:
		fmt.Println("Invalid choice.")
		todo()
	}
}
