package main

func validateName (name string) bool {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		println("Name cannot be empty. Please enter a valid name.")
		return false
	}