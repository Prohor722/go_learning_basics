package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("Doing task:", id)
}

func main() {
	for i:=range 11 {
		// go task(i)
		go func(i int){
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Microsecond*200)
}