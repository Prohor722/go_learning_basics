package main

import (
	"fmt"

	"github.com/Prohor722/go_learning_basics/account"
	"github.com/Prohor722/go_learning_basics/auth"
	"github.com/Prohor722/go_learning_basics/user"
	"github.com/fatih/color"
)

// go mod init github.com/Prohor722/go_learning_basics	//cmd for creating package
//naming by the github project url is convension not mendatory

//go get github.com/fatih/color		//to install package
//go mod tidy 		//to fix basic warning issues fixed

func main() {
	auth.LoginWithCredentials("Prohor", "1234")
	session := auth.GetSession()
	fmt.Println(session)

	user := user.User{
		Email: "user@user.com",
		Name:  "User1",
	}

	fmt.Println(user.Email)
	fmt.Println(user.Name)

	user1 := account.UserAccount{
		Name: "Abu",
		Age:  10,
	}
	user1.CreateAccount()

	fmt.Println(user1.Name)
	fmt.Println(user1.CheckBalance())
	user1.AddAmount(-11.22)
	fmt.Println("After add -11.20")
	fmt.Println(user1.CheckBalance())
	user1.AddAmount(35)
	fmt.Println("After add 35")
	fmt.Println(user1.CheckBalance())
	user1.SubAmount(-40)
	fmt.Println("After sub -40")
	fmt.Println(user1.CheckBalance())
	user1.SubAmount(12)
	fmt.Println("After sub 12")
	fmt.Println(user1.CheckBalance())

	//using installed package
	color.Red(user1.Name)
	color.Yellow(user1.Name)
}
