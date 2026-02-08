package main

type Contact struct {
	Name  string
	Phone string
	Email string
}
var contacts []*Contact
var deletedContacts []*Contact
var options = []string{
	"Add Contact", 
	"View Contacts", 
	"Update Contact", 
	"Delete contact", 
	"View Deleted Contacts", 
	"Exit"}

func main() {
	phoneBookApp()
}