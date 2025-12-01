package main

func main() {
	print("Prohor ")
	println("Banik")

	println("print \\ slash")
	println("print \"double quation\".")
	println("line \nbreak") // \n will not create a new line here
	println("add tab \t here") // \t will not create a tab here

	// Variables()

	showVariables();
}

func Variables(){
	a:= 10
	b:= 20
	sum := a + b
	println("Sum is:", sum)
}
