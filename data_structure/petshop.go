package datastructure 


import (
	"fmt"
)


func registerPet() error {

	for {
		var choice int

		fmt.Println("What pet you want to register?")
		fmt.Println("1-> Dog")
		fmt.Println("2-> Cat")
		fmt.Println("Enter your choice")
		if _, err := fmt.Scanln(&choice);
		err != nil {
			return fmt.Errorf("failed to get choice")
		}

		switch choice {

		case 1:
			fmt.Println("Not ready yet")
			// registerDog()
			continue

		case 2:
			registerCat()
			continue

		case 3:

		default:
			fmt.Println("Exited...")
			return nil
		}
	}
}

