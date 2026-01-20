package main

func findLargestSmallestElement(){
	numbers := []int{34, 12, 5, 67, 23, 89, 2, 45}
	if len(numbers) == 0 {
		return
	}
	largest := numbers[0]
	smallest := numbers[0]
	for _, num := range numbers {
		if num > largest {
			largest = num
		}
		if num < smallest {
			smallest = num
		}
	}
	println("Largest element:", largest)
	println("Smallest element:", smallest)
}