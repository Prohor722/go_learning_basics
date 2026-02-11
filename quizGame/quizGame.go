package main

func quizGame() {
	var name string
	var age uint = 0

	println("Welcome to the Quiz Game!")
	name = getName()
	age = getAge(name)

	println("Hello", name, "age", age, "! Let's start the quiz.")

	if(age < 12) {
		getUserAnswer(quiz.A_questions)
	} else if(age < 15) {
		getUserAnswer(quiz.B_questions)
	} else {
		getUserAnswer(quiz.C_questions)
	}

	
	println("\nThank you for playing, ", name, "!")

}
