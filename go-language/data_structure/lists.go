package datastructure

import "fmt"


func arrayDemo() {
	
	var productNames [4]string
	var printEmptyArray[4]string 

	//basic array printing
	prices := [4]float64{10.99, 9.99, 22.22, 8.77}
	fmt.Println(prices)
	fmt.Println(prices[0])
	fmt.Println(printEmptyArray)

	productNames[2] = "Hello"
	fmt.Println(productNames)
	
	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)
}
