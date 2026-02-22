package main

import (
	"fmt"
	"github.com/Prohor722/go_learning_basics/auth"
	"github.com/Prohor722/go_learning_basics/user"
)

// go mod init github.com/Prohor722/go_learning_basics	//cmd
//naming by the github project url is convension not mendatory

func main(){
	auth.LoginWithCredentials("Prohor","1234")
	session := auth.GetSession()
	fmt.Println(session)

	user := user.User{
		Email: "user@user.com",
		Name: "User1",
	}

	fmt.Println(user.Email)
	fmt.Println(user.Name)
}