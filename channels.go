package main

import (
	"fmt"
	"math/rand/v2"
	"time"
	// "time"
)

//Sending data...
func processNum(numChan chan int){
	for num := range numChan{
		fmt.Println("processing number", num)
		time.Sleep(time.Second*1)
	}
	// fmt.Println("processing number", <-numChan)
}

func channelsExample() {
	numChan := make(chan int)

	go processNum(numChan)

	// numChan <- 5

	for{
		numChan <- rand.IntN(100)
	}

	// time.Sleep(time.Second*2)

	// messageChan := make(chan string)
	// messageChan <- "ping..."	//blocking
	// msg := <-messageChan
	// fmt.Println(msg)
}