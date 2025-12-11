package main

func assignOperator() {
	x := 10
	println("when, x := 10 ",x)
	x += 5	  // x = x + 5
	println("after x += 5 ",x)
	x -= 3   // x = x - 3
	println("after x -= 5 ",x)
	x *= 2   // x = x * 2
	println("after x *= 2 ",x)
	x /= 4   // x = x / 4
	println("after x /= 4 ",x)
	x %= 3   // x = x % 3
	println("after x %= 3 ",x)
	x++;
	println("after x++ ",x)
	x--;
	println("after x-- ",x)


}