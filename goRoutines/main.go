package main

import (
	"fmt"
	"sync"
	// "time"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Doing task:", id)
}

func main() {
	var wg sync.WaitGroup
	for i:=range 11 {
		wg.Add(1)
		go task(i, &wg)
		// go func(i int){
		// 	fmt.Println(i)
		// }(i)
	}

	wg.Wait()

	// time.Sleep(time.Microsecond*200)
}