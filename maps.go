package main

import (
	"fmt"
	"maps"
)

func maps() {
	m := make(map[string]string)
	m1 := map[string]int{
		"one": 1,
		"two": 2,
	}
	m2 := map[string]int{
		"one": 1,
		"two": 2,
	}

	fmt.Println(maps.Equal(m1, m2)) // true

	m["name"] = "Go"
	m["age"] = "15"
	m["type"] = "Programming Language"

	fmt.Println(m["name"])
	fmt.Println(m["type"])
	fmt.Println(m["nonexistent"]) // This will print an empty string

	if m["nonexistent"] == "" {
		fmt.Println("true")
	}

	k,ok := m["nonexistent"] //if k is not needed use _

	if !ok {
		fmt.Println("Key does not exist")
	}
	fmt.Println(k)
}