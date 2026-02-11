package main

var questionsA = []string{
		"What is the capital of France?",
		"What is 2 + 2?",
		"What is the largest planet in our solar system?",
	}

var optionsA = [][]string{
		{"A) Berlin", "B) Madrid", "C) Paris", "D) Rome"},
		{"A) 3", "B) 4", "C) 5", "D) 6"},
		{"A) Earth", "B) Jupiter", "C) Mars", "D) Saturn"},
	}

var correctAnswersA = []string{"C", "B", "B"}

func main() {
	quizGame()
}