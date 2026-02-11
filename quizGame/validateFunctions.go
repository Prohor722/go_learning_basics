package main

import (
	"strings"
)

func validateName (name string) bool {
	name = strings.TrimSpace(name)
	if len(name) < 2 {
		println("Please enter a valid name (at least 2 characters).")
		return false
	}
	return true
}