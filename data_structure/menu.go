package datastructure

import (
	"fmt"
)


func PetshopMenu() error {

	for {

		var choice int

		fmt.Println("---Petshop menu---")
		fmt.Println("1->Register a dog or a cat")
		fmt.Printf("Enter your choice: ")
		if _, err := fmt.Scanln(&choice); err != nil {
			return fmt.Errorf("failed to get the choice")
			// continue
		}
	
		switch choice {
			
		case 1: 
			registerPet()
			continue

		case 2:
			fmt.Println("exited program...")
			return nil

		default:
			fmt.Println("invalid choice")
			continue
		}
	}
}

