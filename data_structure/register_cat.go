package datastructure

import (
	"errors"
	"fmt"
)


var catList []cat

func registerCat() error {

	fmt.Println("Cat registration")
	fmt.Printf("Enter the cat name: ")
	var catName string
	if _, err := fmt.Scanln(&catName);
	err != nil {
		return fmt.Errorf("failed to read name")
	}

	fmt.Printf("Enter the cat breed: ")
	var catBreed string
	fmt.Scanln(&catBreed)

	fmt.Printf("Enter the cat age: ")
	var catAge int
	if _, err := fmt.Scanln(&catAge);
	err != nil {
		return fmt.Errorf("failed to read cat age")
	}

	newCatData, err := newPet(catName, catBreed, catAge)	
	if err != nil {
		return fmt.Errorf("could not save data %w", err)
	}

	catList = append(catList, *newCatData)

	fmt.Println(catList)
	fmt.Println(catList[:1])

	return nil
}

type cat struct {
	name string
	breed string
	age int
}

func newPet(name, breed string, age int) (*cat, error){
	if breed == "" {
		 breed = "stray"
	}

	if name == "" {
		return nil, errors.New("cat name cannot be empty")
	}

	if age <= 0 {
		return nil, errors.New("cat age cannot be 0")
	}

	return &cat{
		name: name,
		breed: breed,
		age: age,
	}, nil
}

