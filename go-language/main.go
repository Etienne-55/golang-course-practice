package main

import (
	"fmt"
	"go-course/projects"
	"go-course/user_account"
	datastructure "go-course/data_structure"
	pricecalculator "go-course/price_calculator"
)


func main() {

for{
var switchChoice int
		fmt.Println("---Golang Projects---")
		fmt.Println("1-> Register acount")
		fmt.Println("2-> Notepad")
		fmt.Println("3-> Register a pet")
		fmt.Println("4-> Price calculator")
		fmt.Println("5-> Tests")
		fmt.Println("6-> End operation")
		fmt.Printf("Choose an operaration: ")
		fmt.Scan(&switchChoice)
		fmt.Printf("Your choice: %d\n", switchChoice)

		switch switchChoice {
		case 1:
			useraccount.RegisterUser()
			continue

		case 2:
			projects.Notepad()
			continue

		case 3:
			datastructure.RegisterPet()
			continue

		case 4:
			pricecalculator.PriceCalculator()
			return
			
		case 5:
			fmt.Println("case used for testing golang features alone")
			// mapsTest()
			structMapArray()
			return

		case 6:
			fmt.Println("Exited program")
			return

		default:
			fmt.Println("error. wrong input ")
		}
	}
}



func mapsTest() error {

	var userList []map[string]int

	bigMap := make(map[string]int)
	
	fmt.Println("Enter your name: ")
	var name string
	if _, err := fmt.Scanln(&name);
	err != nil {
		return fmt.Errorf("cant register name")
	}

	fmt.Println("Enter your age: ")
	var age int 
	if _, err := fmt.Scanln(&age);
	err != nil {
		return fmt.Errorf("cant register age")
	}

	bigMap[name] = age

	user := bigMap
	userList = append(userList, user)


	fmt.Println(bigMap)
	fmt.Println(userList)

	return nil
}

	type user struct {
		name string
		email string
	}

func structMapArray() error {

	userMap := make(map[int]user)

	var userList []map[int]user

	fmt.Println("Enter your name: ")
	var name string
	if _, err := fmt.Scanln(&name);
	err != nil {
		return fmt.Errorf("error %w", err)
	}

	fmt.Println("Enter your email: ")
	var email string
	if _, err := fmt.Scanln(&email);
	err != nil {
		return fmt.Errorf("error %w", err)
	}

	userData, nil := newUser(name, email)

	fmt.Println("Enter your age: ")
	var age int 
	if _, err := fmt.Scanln(&age);
	err != nil {
		return fmt.Errorf("error %w", err)
	}

	userMap[age] = userData

	userList = append(userList, userMap)


	fmt.Println(userList)

	return nil
} 

func newUser(name, email string ) (user, error) {
	if name == "" {
		return user{}, fmt.Errorf("name cant be blank")
	}

	if email == "" {
		return user{}, fmt.Errorf("email cant be blank")
	}

	return user {
		name: name,
		email: email,
	}, nil
}

