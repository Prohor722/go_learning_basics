package main

import "fmt"

func maps() {
	m := make(map[string]string)

	m["name"] = "Go"
	m["type"] = "Programming Language"

	fmt.Println(m["name"])

}