package main

import (
	"strings"
)

func nameValidation(name string) bool {
	if len(name) < 2 {
		return false
	}
	return true
}

func ageValidation(age int) bool {
	if age < 0 || age > 120 {
		return false
	}
	return true
}

func validation(validationType string, value interface{}) bool {
	switch validationType {
	case "name":
		name, ok := value.(string)
		if !ok || len(name) < 2 {
			return false
		}
		return true
	case "age":
		age, ok := value.(int)
		if !ok || age < 0 || age > 120 {
			return false
		}
		return true
	case "email":
		email, ok := value.(string)
		if !ok || len(email) < 5 || !strings.Contains(email, "@") {
			return false
		}
		return true
	default:
		return false
	}
}