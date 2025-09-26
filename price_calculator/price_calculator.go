package pricecalculator

import (
	"os"
	"fmt"
	"bufio"
	"errors"
)


func calculeAndWrite() {
	
	fmt.Println("")	


}

type product struct {
	originalPrice int
	taxesApplied int
	newPrice int
}


func newProduct(originalPrice, taxesApplied, newPrice int) (*product, error){
	if originalPrice <= 0 {
		return nil, errors.New("price cant be less or equal to 0")
	}	

	return &product{
		originalPrice: originalPrice,
		taxesApplied: taxesApplied,
		newPrice: newPrice,
	}, nil
}

func (c *product) OriginalPrice() int { return c.originalPrice }
func (c *product) TaxesApplied() int { return c.taxesApplied }
func (c *product) NewPrice() int { return c.newPrice }


func calculatePrice() (Product, error)  {
	fmt.Printf("Enter  the value of the product: ")
	var productPrice int
	if _, err := fmt.Scan(&productPrice);
	err != nil {
		return nil, fmt.Errorf("error getting price input")
	}

	fmt.Printf("Enter  the value of the product: ")
	var taxValue int
	if _, err := fmt.Scan(&taxValue);
	err != nil {
		return nil, fmt.Errorf("error getting price input")
	}

	var finalValue int
	finalValue = productPrice * taxValue


	return newProduct(productPrice, taxValue, finalValue)

}


// func savePriceToFile() error {
// 	fmt.Pr
//
//
//
//
//
// }

func writeFile() error {
	
	fmt.Print("Enter the note name: ")
	var fileName string
	if _, err := fmt.Scanln(&fileName); err != nil {
		return fmt.Errorf("failed to read filename: %w", err)
	}

	fmt.Print("Enter the note: ")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return fmt.Errorf("failed to read note content")
	}

	formattedFileName := fileName + ".txt"
	if err := os.WriteFile(formattedFileName, []byte(scanner.Text()), 0644);
	err != nil {
		return  fmt.Errorf("Failed to write file: %w", err)
	}

	fmt.Printf("Note saved to %s\n", formattedFileName)
	return nil
}

