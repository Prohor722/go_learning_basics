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

func emailSender(emailChan <-chan string, done chan<- bool){
	defer func(){ done<-true}()
	for email:=range emailChan{
		fmt.Println("sending email to: ",email)
		time.Sleep(time.Second)
	}
}

func channelsExample() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func(){
		chan1 <- 1
	}()

	go func(){
		chan2 <- "pong.."
	}()

	for range 2{
		select {
			case chan1Val := <-chan1:
				fmt.Println("receiving data from chan1: ",chan1Val)
			case chan1Val := <-chan2:
				fmt.Println("receiving data from chan2: ",chan1Val)
		}
	}
	// emailChan := make(chan string, 10)
	// done := make(chan bool)
	// go emailSender(emailChan, done)
	// for i:=range 100{
	// 	emailChan <- fmt.Sprintf("mail%d@gmail.com",i)	
	// }
	// fmt.Println("done sending.")
	// close(emailChan)
	// <-done

	// done := make(chan bool)
	// go task(done)
	// res:= <- done
	// fmt.Println(res)

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