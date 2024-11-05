//Get word from user the print the word in alphabetic order.
package main

import (
	"fmt"
	"sort"
)

func main() {

	var word string
	fmt.Println("Enter any word: ")
	fmt.Scanln(&word)

	runes := []rune(word)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})


	sorted_word := string(runes)

	fmt.Println("Sorted word:", sorted_word)
}
