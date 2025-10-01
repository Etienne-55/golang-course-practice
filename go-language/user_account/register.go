package useraccount

import (
	"os"
	"fmt"
	"encoding/json"
)


func RegisterUser() error {

	fmt.Print("Enter your email: ")
	var email string
	if _, err := fmt.Scanln(&email); err != nil {
		return fmt.Errorf("failed to read email")
	}
	
	fmt.Print("Enter password: ")
	var password string
	if _, err := fmt.Scanln(&password); err != nil {
		return fmt.Errorf("failed to read password")
	}

	newUserData, _ := NewUser(email, password)	

	jsonData, err := json.Marshal(newUserData) 
	if err != nil {
		return err
	}

	formattedTitle := email + ".json"
	os.WriteFile(formattedTitle,[]byte(jsonData), 0644)

	return nil
}

