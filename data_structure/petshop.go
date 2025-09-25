package datastructure 

import (
	"fmt"
)


var PetList []Pet

func RegisterPet() error {

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
			if err := register(registerDog);
			nil != err {
				fmt.Println(err)
			}
			continue

		case 2:
			if err := register(registerCat);
			nil != err {
				fmt.Println(err)
			}
			continue

		case 3:

		default:
			fmt.Println("Exited...")
			return nil
		}
	}
}

type Pet interface {
	AnimalType() string
	Name() string
	Breed() string
	Age() int
}


func register(readFunc func() (Pet, error)) error {
	pet, err := readFunc()
	if err != nil {
		return fmt.Errorf("registration failed %w", err)
	}
	PetList = append(PetList, pet)

	for i, p := range PetList {
		fmt.Printf("%d. ->%s %s (%s), Age: %d\n", i+1, p.AnimalType(), p.Name(), p.Breed(), p.Age())
	}

	return nil
}

