package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userfirstName := getUserData("Please enter your first name 2: ")
	userlastName := getUserData("Please enter your last name: ")
	userbirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	person, error := user.New(userfirstName, userlastName, userbirthdate)
	admin := user.NewAdmin("prezes.sikora@gmail.com", "secret")
	admin.User.OutputUser()
	if error != nil {
		fmt.Println(error)
		return
	}

	// ... do something awesome with that gathered data!

	person.OutputUser()
	person.ClearUserName()
	person.OutputUser()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
