package datastructure

import (
	"fmt"
)


func TestPointers() {
	
	var age = 30

	agePointer := &age

	adultAge := getAdultYears(agePointer)

	fmt.Println(adultAge)
}


func getAdultYears(age *int) int {

	return *age - 18
}
