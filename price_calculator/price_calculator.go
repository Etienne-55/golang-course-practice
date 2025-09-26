package pricecalculator

import (
	"os"
	"fmt"
	"errors"
	"encoding/json"
)


func PriceCalculatorMain() {
	productInfo, _ := calculatePrice()
	savePriceToFile(productInfo)
} 

func calculatePrice() (*product, error)  {
	fmt.Printf("Enter the value of the product: ")
	var productPrice float64
	if _, err := fmt.Scan(&productPrice);
	err != nil {
		return nil, fmt.Errorf("error getting price input")
	}

	fmt.Printf("Enter the value of the tax: ")
	var taxValue float64
	if _, err := fmt.Scan(&taxValue);
	err != nil {
		return nil, fmt.Errorf("error getting price input")
	}

	realTaxValue := taxValue / 100
	fmt.Println(realTaxValue)

	finalValue := productPrice * (1 + realTaxValue)
	fmt.Printf("The final price is %2.f\n", finalValue)

	return newProduct(productPrice, taxValue, finalValue)
}

func savePriceToFile(product *product) error {
	jsonData, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("failed to tranform data to json")
	}

	if err := os.WriteFile("price_after_taxes.json", []byte(jsonData), 0644);
	err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	fmt.Println("price saved to files")
	return nil
}

type product struct {
	OriginalPrice float64
	TaxesApplied float64
	NewPrice float64 
}

func newProduct(originalPrice, taxesApplied, newPrice float64) (*product, error){
	if originalPrice <= 0 {
		return nil, errors.New("price cant be less or equal to 0")
	}	

	return &product{
		OriginalPrice: originalPrice,
		TaxesApplied: taxesApplied,
		NewPrice: newPrice,
	}, nil
}

