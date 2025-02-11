package main

import (
	"errors"
	"fmt"

	"sikora.com/notes/notes"
)

func getUserInupt(prompt string) (string, error) {
	var value string
	fmt.Print(prompt)
	fmt.Scanln(&value)
	if value == "" {
		return "", errors.New("invalid inupt")
	}
	return value, nil
}

func main() {
	var i interface{}
	i = 10
	value, ok := i.(int)
	fmt.Println(value)
	fmt.Println(ok)

	i = "sssss"
	value2, ok2 := i.(string)
	fmt.Println(value2)
	fmt.Println(ok2)

	title, error := getUserInupt("Note title: ")
	if error != nil {
		fmt.Print(error)
		return
	}

	content, error := getUserInupt("Note content: ")
	if error != nil {
		fmt.Print(error)
		return
	}

	note := notes.New(title, content)
	note.Display()
	err := note.Save()
	if err != nil {
		fmt.Printf("Error saving files : %v", error)
		return
	} else {
		fmt.Println("Saving Sucess.!")
	}
}
