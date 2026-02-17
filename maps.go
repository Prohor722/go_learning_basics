package main

import "fmt"

func maps() {
	m := make(map[string]string)

	m["name"] = "Go"
	m["type"] = "Programming Language"

	fmt.Println(m["name"])
	fmt.Println(m["type"])
	fmt.Println(m["nonexistent"]) // This will print an empty string
}