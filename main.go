package main

import (
	"fmt"
	datastructure "go-course/data_structure"
	"go-course/projects"
	"go-course/user_account"
)


func main() {

for{
var switchChoice int
		fmt.Println("---Golang Projects---")
		fmt.Println("1-> Register acount")
		fmt.Println("2-> Notepad")
		fmt.Println("3-> Register a pet")
		fmt.Println("3-> End operation")
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
			fmt.Println("Exited program")
			return
			
		default:
			fmt.Println("error. wrong input ")
		}
	}
}

