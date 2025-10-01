package main

import "fmt"

func PointersLesson() {
	age := 32
	var agePointer *int

	agePointer = &age

	fmt.Println(agePointer)

	adultYears := getAdultYears(age)
	fmt.Println(adultYears)
} 

func getAdultYears(age int) int {
	return age - 18
}
