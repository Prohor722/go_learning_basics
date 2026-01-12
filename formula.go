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

func GeometricProgression() {
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

func fibonacciSeries() {
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

func sumOfFirstNNumbers() {
	var n int
	print("Enter a positive integer:")
	fmt.Scanln(&n)
	sum := n * (n + 1) / 2
	println("Sum of the first", n, "natural numbers is:", sum)
}

func SumOfSquares() {
	var n int
	print("Enter a positive integer:")
	fmt.Scanln(&n)
	sum := n * (n + 1) * (2*n + 1) / 6
	println("Sum of squares of the first", n, "natural numbers is:", sum)
}

func sumOfCubes() {
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

func permutation() {
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

func combination() {
	var n, r int
	print("Enter n (total items):")
	fmt.Scanln(&n)

	print("Enter r (items to choose):")
	fmt.Scanln(&r)
	numerator := factorial(n)
	denominator := factorial(r) * factorial(n-r)
	result := numerator / denominator
	println("Number of combinations (C(n,r)) is:", result)
}

func power() {
	var base, exponent int
	print("Enter the base:")
	fmt.Scanln(&base)
	print("Enter the exponent:")
	fmt.Scanln(&exponent)

	result := int(math.Pow(float64(base), float64(exponent)))
	println(base, "raised to the power of", exponent, "is:", result)
}

func logIdentity() {
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

func changeOfBase() {
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

func additionFormula() { //(a+b)%m=((a%m)+(b%m))%m
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

func multiplicationFormula() { //(a*b)%m=((a%m)*(b%m))%m
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

func moduloInverse() {
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

func GCD() {
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

func LCM() {
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

func primeNumberFormulas() {
	var n int
	print("Enter a positive integer:")
	fmt.Scanln(&n)
	isPrime := true
	if n <= 1 {
		isPrime = false
	} else {
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
	}
	if isPrime {
		println(n, "is a prime number.")
	} else {
		println(n, "is not a prime number.")
	}
}

func moduloExponentiation() {
	var base, exponent, mod int
	print("Enter the base:")
	fmt.Scanln(&base)
	print("Enter the exponent:")
	fmt.Scanln(&exponent)
	print("Enter the modulus:")
	fmt.Scanln(&mod)
	result := 1
	base = base % mod
	for exponent > 0 {
		if exponent%2 == 1 {
			result = (result * base) % mod
		}
		exponent = exponent >> 1
		base = (base * base) % mod
	}
	println("Result of (base^exponent) % mod is:", result)
}

func chineseRemainderTheorem() {
	var a1, a2, m1, m2 int
	print("Enter a1:")
	fmt.Scanln(&a1)
	print("Enter m1:")
	fmt.Scanln(&m1)
	print("Enter a2:")
	fmt.Scanln(&a2)
	print("Enter m2:")
	fmt.Scanln(&m2)
	M := m1 * m2
	M1 := M / m1
	M2 := M / m2
	inv1 := 1
	for i := 1; i < m1; i++ {
		if (M1*i)%m1 == 1 {
			inv1 = i
			break
		}
	}

	inv2 := 1
	for i := 1; i < m2; i++ {
		if (M2*i)%m2 == 1 {
			inv2 = i
			break
		}
	}
	println("The solution to the system of congruences is:", (a1*M1*inv1+a2*M2*inv2)%M)
}

func countSetBits(n int) int {
	count := 0
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

func bitManipulation() {
	var n int
	print("Enter an integer:")
	fmt.Scanln(&n)
	println("Number of set bits in", n, "is:", countSetBits(n))
}

func BitManipulationCheckEven(){
	var n int
	print("Enter an integer:")
	fmt.Scanln(&n)
	if n&1 == 0 {
		println(n, "is even.")
	} else {
		println(n, "is odd.")
	}
}

func BitManipulationSwap(){
	var a, b int
	print("Enter value of a:")
	fmt.Scanln(&a)
	print("Enter value of b:")
	fmt.Scanln(&b)
	println("Before swapping: a =", a, ", b =", b)
	a = a ^ b
	b = a ^ b
	a = a ^ b
	println("After swapping: a =", a, ", b =", b)
}

func BitManipulationPowerOfTwo(){
	var n int
	print("Enter an integer:")
	fmt.Scanln(&n)
	if n > 0 && (n&(n-1)) == 0 {
		println(n, "is a power of two.")
	}else {
		println(n, "is not a power of two.")
	}
}

func bitManipulationClearBits(){
	var n, k int
	print("Enter an integer:")
	fmt.Scanln(&n)
	print("Enter number of least significant bits to clear (k):")
	fmt.Scanln(&k)
	mask := ^((1 << k) - 1)
	result := n & mask
	println("Result after clearing", k, "least significant bits of", n, "is:", result)
}

func bitManipulationToggleBits(){
	var n, k int
	print("Enter an integer:")
	fmt.Scanln(&n)
	print("Enter number of least significant bits to toggle (k):")
	fmt.Scanln(&k)
	mask := (1 << k) - 1
	result := n ^ mask
	println("Result after toggling", k, "least significant bits of", n, "is:", result)
}

func bitManipulationTurnOffRightmostSetBit(){
	var n int
	print("Enter an integer:")
	fmt.Scanln(&n)
	result := n & (n - 1)
	println("Result after turning off the rightmost set bit of", n, "is:", result)
}