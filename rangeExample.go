package main

//range is used to iterate over elements in a variety of data structures, such as arrays, slices, maps, and strings. It provides a convenient way to loop through these data structures without having to manage the index or key manually.
func rangeExample() {
	nums := []int{1, 2, 3, 4, 5}
	for index, value := range nums {
		println("Index:", index, "Value:", value)
	}

	for value := range nums {
		println("Value:", value)
	}

	for i,c := range "Hello" {
		println("Index:", i, "Character:", string(c))
	}
}