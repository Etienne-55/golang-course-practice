package pricecalculator

import (
	"fmt"
)


func PriceCalculator() error {

	for {
		var choice int

		fmt.Println("--Price calculator--")
		fmt.Println("1->Calculate and save price after taxes")
		fmt.Printf("Enter your choice: ")
		if _, err := fmt.Scan(&choice);
		err != nil {
			return fmt.Errorf("failed to read choice")
		}

		switch choice {

		case 1:
			PriceCalculatorMain()
			continue

		case 2:
			fmt.Println("exiting program...")
			return nil

		default:
			fmt.Println("error, wrong input")
			continue
		}
	}
}


