package main

func allPrimesToN(n int) []int {
	primes := []int{}
	for i := 2; i <= n; i++ {
		isPrime := true
		