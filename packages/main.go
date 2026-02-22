package main

import "github.com/Prohor722/go_learning_basics/auth"

// go mod init github.com/Prohor722/go_learning_basics	//cmd
//naming by the github project url is convension not mendatory

func main(){
	auth.LoginWithCredentials("Prohor","1234")
}