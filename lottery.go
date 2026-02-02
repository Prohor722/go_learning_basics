package main

import "fmt"

func drawLottery(winningNumbers []int, userNumbers []int) {
	matchingNumbers := []int{}
	for _, userNum := range userNumbers {
		for _, winNum := range winningNumbers {
			if userNum == winNum {