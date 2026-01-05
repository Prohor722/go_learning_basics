package main

import (
	"fmt"
	"math"
)

func sequenceFormula() {
	var n int
	print("Enter the value of n:")
	fmt.Scanln(&n)

	result := n * (n + 1) / 2
	println("The sum of the first", n, "natural numbers is:", result)
}

func arithmeticProgression() {
	var a, d, n int
	print("Enter the first term (a):")
	fmt.Scanln(&a)
	print("Enter the common difference (d):")
	fmt.Scanln(&d)
	print("Enter the number of terms (n):")
	fmt.Scanln(&n)

	lastTerm := a + (n-1)*d
	sum := n * (a + lastTerm) / 2
	println("The sum of the arithmetic progression is:", sum)
}

func GeometricProgression(){
	var a, r, n int
	print("Enter the first term (a):")
	fmt.Scanln(&a)
	print("Enter the common ratio (r):")
	fmt.Scanln(&r)
	print("Enter the number of terms (n):")
	fmt.Scanln(&n)
	sum := a * (int(math.Pow(float64(r), float64(n))) - 1) / (r - 1)
	println("The sum of the geometric progression is:", sum)
}

func fibonacciSeries(){
	var n int
	print("Enter the number of terms:")
	fmt.Scanln(&n)
	a, b := 0, 1
	println("Fibonacci Series:")
	for i := 0; i < n; i++ {
		print(a, " ")
		a, b = b, a+b
	}
}

func sumOfFirstNNumbers(){
	var n int
	print("Enter a positive integer:")
	fmt.Scanln(&n)
	sum := n * (n + 1) / 2
	println("Sum of the first", n, "natural numbers is:", sum)
}

func SumOfSquares(){
	var n int
	print("Enter a positive integer:")
	fmt.Scanln(&n)
	sum := n * (n + 1) * (2*n + 1) / 6
	println("Sum of squares of the first", n, "natural numbers is:", sum)
}

func sumOfCubes(){
	var n int
	print("Enter a positive integer:")
	fmt.Scanln(&n)
	sum := (n * (n + 1) / 2) * (n * (n + 1) / 2)
	println("Sum of cubes of the first", n, "natural numbers is:", sum)
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func permutation(){
	var n, r int
	print("Enter n (total items):")
	fmt.Scanln(&n)
	print("Enter r (items to arrange):")
	fmt.Scanln(&r)
	numerator := factorial(n)
	denominator := factorial(n - r)
	result := numerator / denominator
	println("Number of permutations (P(n,r)) is:", result)
}

func combination(){
	var n, r int
	print("Enter n (total items):")
	fmt.Scanln(&n)

	print("Enter r (items to choose):")
	fmt.Scanln(&r)
	numerator := factorial(n)
	denominator := factorial(r) * factorial(n - r)
	result := numerator / denominator
	println("Number of combinations (C(n,r)) is:", result)
}

func power(){
	var base, exponent int
	print("Enter the base:")
	fmt.Scanln(&base)
	print("Enter the exponent:")
	fmt.Scanln(&exponent)

	result := int(math.Pow(float64(base), float64(exponent)))
	println(base, "raised to the power of", exponent, "is:", result)
}

func logIdentity(){
	var x float64
	print("Enter a positive number:")
	fmt.Scanln(&x)
	if x <= 0 {
		println("Logarithm is undefined for non-positive numbers.")
		return
	}
	result := math.Log(x * x)
	println("log(x^2) is:", result)
}

func changeOfBase(){
	var x, base float64
	print("Enter a positive number (x):")
	fmt.Scanln(&x)
	print("Enter the base (b):")
	fmt.Scanln(&base)
	if x <= 0 || base <= 1 {
		println("Invalid input for logarithm.")
		return
	}
	result := math.Log(x) / math.Log(base)
	println("log_b(x) is:", result)
}

func additionFormula(){		//(a+b)%m=((a%m)+(b%m))%m
	var a, b, m int
	print("Enter a:")
	fmt.Scanln(&a)
	print("Enter b:")
	fmt.Scanln(&b)
	print("Enter m:")
	fmt.Scanln(&m)
	result := ((a % m) + (b % m)) % m
	println("(a+b)%m is:", result)
}

func multiplicationFormula(){		//(a*b)%m=((a%m)*(b%m))%m
	var a, b, m int
	print("Enter a:")
	fmt.Scanln(&a)
	print("Enter b:")
	fmt.Scanln(&b)
	print("Enter m:")
	fmt.Scanln(&m)
	result := ((a % m) * (b % m)) % m
	println("(a*b)%m is:", result)
}

// func logIdentity(){
// 	var x float64
// 	print("Enter a positive number:")
// 	fmt.Scanln(&x)
// 	if x <= 0 {
// 		println("Logarithm is undefined for non-positive numbers.")
// 		return
// 	}
// 	result := math.Log(x * x)
// 	println("log(x^2) is:", result)
// }

func moduloInverse(){
	var a, m int
	print("Enter a:")
	fmt.Scanln(&a)
	print("Enter m (prime):")
	fmt.Scanln(&m)
	inverse := 1
	for i := 1; i < m; i++ {
		if (a*i)%m == 1 {
			inverse = i
			break
		}
	}
	println("The modulo inverse of", a, "mod", m, "is:", inverse)
}

func GCD(){
	var a, b int
	print("Enter a:")
	fmt.Scanln(&a)
	print("Enter b:")
	fmt.Scanln(&b)
	for b != 0 {
		a, b = b, a%b
	}
	println("The GCD is:", a)
}

func LCM(){
	var a, b int
	print("Enter a:")
	fmt.Scanln(&a)
	print("Enter b:")
	fmt.Scanln(&b)
	gcd := a
	bTemp := b

	for bTemp != 0 {
		gcd, bTemp = bTemp, gcd%bTemp
	}

	lcm := (a * b) / gcd
	println("The LCM is:", lcm)
}
