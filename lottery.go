package main

import "fmt"

func drawLottery(winningNumbers []int, userNumbers []int) {
	matchingNumbers := []int{}
	for _, userNum := range userNumbers {
		for _, winNum := range winningNumbers {
			if userNum == winNum {
				matchingNumbers = append(matchingNumbers, userNum)
			}
		}
	}
	if len(matchingNumbers) == 0 {
		fmt.Println("No matching numbers. Better luck next time!")
	}