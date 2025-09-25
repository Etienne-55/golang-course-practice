package datastructure

import (
	"os"
	"fmt"
	"bufio"
	"errors"
)

func registerCat() (Pet, error) {
	
	animalCategory := "cat"

	fmt.Println("Cat registration")
	fmt.Printf("Enter the cat name: ")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return nil, fmt.Errorf("failed to read name")
	}
	catName := scanner.Text()

	fmt.Printf("Enter the cat breed: ")
	var catBreed string
	fmt.Scanln(&catBreed)

	fmt.Printf("Enter the cat age: ")
	var catAge int
	if _, err := fmt.Scanln(&catAge);
	err != nil {
		return nil, fmt.Errorf("failed to read cat age")
	}

	return newCat(animalCategory, catName, catBreed, catAge)
}
type cat struct {
	animalType string
	name string
	breed string
	age int
}

func (c *cat) AnimalType() string { return c.animalType }
func (c *cat) Name() string { return c.name }
func (c *cat) Breed() string { return c.breed}
func (c *cat) Age() int { return c.age}

func newCat(animalType, name, breed string, age int) (*cat, error){
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
		animalType: animalType,
		name: name,
		breed: breed,
		age: age,
	}, nil
}

