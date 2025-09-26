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


// type tranformFn func(int) int
//
// func numbers() {
// 	numbers := []int{1, 3, 45}
// 	doubled := transformNumbers(&numbers, double)
// 	tripled := transformNumbers(&numbers, triple)
//
// 	fmt.Println(doubled)
// 	fmt.Println(tripled)
// }
//
//
// func transformNumbers(numbers *[]int, transform tranformFn)  []int {
// 	dNumbers := []int{}
//
// 	for _, val := range *numbers {
// 		dNumbers = append(dNumbers, transform(val))
//
// 	}
// 	return dNumbers
// }
//
// // func getTransformetFunction(numbers *[]int) tranformFn {
// // 	if (*numbers)[0] == 1 {
// // 	return double
// // 	} else {
// // 	return  triple
// // }
//
// func double(number int) int {
// 	return number * 2
// }
//
// func triple(number int) int {
// 	return number * 3
// }

