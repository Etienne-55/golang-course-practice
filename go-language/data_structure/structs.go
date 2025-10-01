package datastructure

import (
	"fmt"
	"time"
	"errors"
)


type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}


func StructMain() {
	userFirstName := getUserData("enter your first name")
	userLastName := getUserData("enter your last name")
	userBirthDate := getUserData("enter your birthday date")

	appUser, err := newUser(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.outputUserData()
	appUser.clearUserName()

	admin := NewAdmin("test", "test")

	admin.clearUserName()
}

func (u User) outputUserData() {
	fmt.Printf(u.firstName, u.lastName, u.birthDate)
}

func (u User) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func getUserData(promptText string) string {
	fmt.Println(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

func newUser(firstName, lastName, birthDate string) (*User, error){

	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last and birthdate are required")
	}
	
	return &User{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}


type Admin struct {
	email string
	password string
	User
}

func NewAdmin(email, password string) Admin {
	return Admin {
		email: email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName: "Admin",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}
}

