package main

import "time"

func multiConditionSwitch() {
	switch time.Now().Weekday() {
	case time.Friday, time.Saturday:
		println("It's a weekend")
	default:
		println("It's a weekday")
	}
}