package main

func conditions(){
	a := 10
	b := 20
	b -= 10
	a += 10

	if a < b {
		println("a is less than b")
	}else if a == b {
		println("a is equal to b")
	}else {
		println("a is greater than b")
	}
}