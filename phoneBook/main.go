package main

type Contact struct {
	Name  string
	Phone string
	Email string
}
var contacts []*Contact
var deletedContacts []*Contact

func main() {
	phoneBookApp()
}