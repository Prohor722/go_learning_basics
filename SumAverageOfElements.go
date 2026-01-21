package main

func sumAndAverageOfElements() {
	numbers := []int{10, 20, 30, 40, 50}
	if len(numbers) == 0 {
		return
	}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	average := float64(sum) / float64(len(numbers))