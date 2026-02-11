package main

type Question struct {
	question string
	options []string
	correctAnswer string
}

type Quiz struct {
	A_questions []Question
	B_questions []Question
	C_questions []Question
}

var quiz = Quiz{
	A_questions: []Question{
			{
				question: "What is the capital of France?",
				options: []string{"A) Berlin", "B) Madrid", "C) Paris", "D) Rome"},
				correctAnswer: "C",
			},
			{
				question: "What is 2 + 2?",
				options: []string{"A) 3", "B) 4", "C) 5", "D) 6"},
				correctAnswer: "B",
			},
			{
				question: "What is the largest planet in our solar system?",
				options: []string{"A) Earth", "B) Jupiter", "C) Mars", "D) Saturn"},
				correctAnswer: "B",
			},
		},
	B_questions: []Question{
			{
				question: "What is the chemical symbol for water?",
				options: []string{"A) H2O", "B) CO2", "C) O2", "D) NaCl"},
				correctAnswer: "A",
			},
			{
				question: "Who wrote 'Romeo and Juliet'?",
				options: []string{"A) Charles Dickens", "B) William Shakespeare", "C) Mark Twain", "D) Jane Austen"},
				correctAnswer: "B",
			},
			{
				question: "What is the speed of light?",
				options: []string{"A) 300,000 km/s", "B) 150,000 km/s", "C) 450,000 km/s", "D) 600,000 km/s"},
				correctAnswer: "A",
			},
		},
	C_questions: []Question{
			{
				question: "What is the derivative of sin(x)?",
				options: []string{"A) cos(x)", "B) -cos(x)", "C) sin(x)", "D) -sin(x)"},
				correctAnswer: "A",
			},
			{
				question: "What is the capital of Iceland?",	
				options: []string{"A) Reykjavik", "B) Oslo", "C) Helsinki", "D) Copenhagen"},
				correctAnswer: "A",
			},
			{
				question: "What is the largest mammal?",
				options: []string{"A) Elephant", "B) Blue Whale", "C) Giraffe", "D) Hippopotamus"},
				correctAnswer: "B",
			},
		},
}

func main() {
	quizGame()
}