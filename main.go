package main

import (
	"fmt"
	"go-course/projects"
)


func main() {

for{
var switchChoice int
		fmt.Println("---Golang Projects---")
		fmt.Println("1-> Notepad")
		fmt.Println("2-> End operation")
		fmt.Printf("Choose an operaration:")
		fmt.Scan(&switchChoice)
		fmt.Printf("Your choice: %d\n", switchChoice)

		switch switchChoice {
		case 1:
			projects.Notepad()
			continue

		case 2:
			fmt.Println("Exited loop")
			return

		default:
			fmt.Println("Exited program")
		}
	}
}

