package main

import (
	"fmt"
	"strings"
)

type Movie struct {
    Name            string
    Price           float64
    AvailableTickets int
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