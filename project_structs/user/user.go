package user

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User     user
}

func NewAdmin(email, password string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User: user{
			firstName: "admin name string",
			lastName:  "admin lastName string",
			birthdate: "string",
			createdAt: time.Now(),
		},
	}
}

func New(userfirstName string, userlastName string, userbirthdate string) (*user, error) {

	if len(userfirstName) < 4 {
		return nil, errors.New("First name should > 3")
	}

	return &user{
		firstName: userfirstName,
		lastName:  userlastName,
		birthdate: userbirthdate,
		createdAt: time.Now(),
	}, nil
}

func (u user) OutputUser() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *user) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}
