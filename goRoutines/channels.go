package main

import "fmt"

func channelsExample() {
	messageChan := make(chan string)

	messageChan <- "ping..."

	msg := <-messageChan

	fmt.Println(msg)
}