package datastructure

import (
	"os"
	"fmt"
	"bufio"
	"errors"
)


func registerDog()  (Pet, error) {
	animalCategory := "dog"

	fmt.Println("dog registration")
	fmt.Printf("Enter the dog name: ")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return nil, fmt.Errorf("failed to read name")
	}
	dogName := scanner.Text()

	fmt.Printf("Enter the dog breed: ")
	var dogBreed string
	fmt.Scanln(&dogBreed)

	fmt.Printf("Enter the dog age: ")
	var dogAge int
	if _, err := fmt.Scanln(&dogAge);
	err != nil {
		return nil, fmt.Errorf("failed to read dog age")
	}
	return newDog(animalCategory, dogName, dogBreed, dogAge)
}

type dog struct {
	animalType string
	name string
	breed string
	age int
}

func (c *dog) AnimalType() string { return c.animalType }
func (c *dog) Name() string { return c.name }
func (c *dog) Breed() string { return c.breed}
func (c *dog) Age() int { return c.age}

func newDog(animalType, name, breed string, age int) (*dog, error) {
	if breed == "" {
		breed = "stray"
	}

	if name == "" {
		return nil, errors.New("dog name cannot be empty")
	}

	if age <= 0 {
		return nil, errors.New("dog age cannot be 0")
	}

	return &dog{
		animalType: animalType,
		name: name,
		breed: breed,
		age: age,
	}, nil
}

