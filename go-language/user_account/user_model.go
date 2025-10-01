package useraccount

import "errors"


type user struct {
	Email string `json:"email"`
	Password string `json:"password"`
}


func NewUser(email, password string) (*user, error){

	if email == "" || password == "" {
		return nil, errors.New("email and password field are required.")
	}

	return &user{
		Email: email,
		Password: password,
	}, nil
}

