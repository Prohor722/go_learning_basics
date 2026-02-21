package main

import (
	"fmt"
	// "math/rand/v2"
	"time"
)

//Sending data...
func processNum(numChan chan int){
	for num := range numChan{
		fmt.Println("processing number", num)
		time.Sleep(time.Second*1)
	}
	// fmt.Println("processing number", <-numChan)
}

func sum(result chan int, num1 int, num2 int){
	numResult := num1 + num2
	result <- numResult
}

func task(done chan bool){
	defer func() {done <- true}()
	fmt.Println("Processing....")
}

func channelsExample() {
	done := make(chan bool)
	go task(done)
	res:= <- done
	fmt.Println(res)

	// result := make(chan int)
	// go sum(result,4,5)
	// res := <- result	//blocking
	// fmt.Println(res)

	// numChan := make(chan int)
	// go processNum(numChan)
	// // numChan <- 5
	// for{
	// 	numChan <- rand.IntN(100)
	// }

	// time.Sleep(time.Second*2)

	// messageChan := make(chan string)
	// messageChan <- "ping..."	//blocking
	// msg := <-messageChan
	// fmt.Println(msg)
}