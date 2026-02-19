package main

func main() {
	// showVariables();
	// Prints();
	// formatVariables();
	// Tests();
	// stingPrints();
	// calculator();

	// assignOperator();
	// logicalOperator();
	
	// conditions();
	// loops();
	// bitwise();
	// evenOdd();
	// largestNumber();
	// gradeCal();
	// bmiCal();
	// switchCase()
	// switchCaseExample2()
	// switchCaseExample3()
	// switchCaseExample4()
	// loopExmple()
	// sequenceFormula()
	// arithmeticProgression()
	// GeometricProgression()
	// fibonacciSeries()
	// sumOfFirstNNumbers()
	// SumOfSquares();
	// sumOfCubes()
	// permutation()
	// combination()
	// power()
	// logIdentity()
	// changeOfBase()
	// additionFormula()
	// multiplicationFormula()
	// quizGame()
	// factorialUsingRecursion()
	// todo()
	// variableActAsConstant()
	// ticketBookingApp()
	// pointerExample()
	// variadicFunction()
	// mapsFunction()
	// clousersExample()
	// pointersExample()
	strucksExamples()
}


func Tests() {
	a:= 111
	b:= 222
	_ = a + b
	sum := a + b

	print(sum)
}

func multiReturn() (int, int) {
	a := 10
	b := 20
	return a, b
}

func printMultiReturn() {
	var arr [2]int
	arr[0], arr[1] = multiReturn()	//in array
	print(arr[0], arr[1])

	a, b := multiReturn()	//in variables
	print(a, b)
}

