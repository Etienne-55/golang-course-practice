package main

import "fmt"


func main() {
	var menuOption int

	for menuOption != 2 {

		fmt.Println("1->Start game")
		fmt.Println("2->End game")
		fmt.Printf("Choose and option: ")
		fmt.Scan(&menuOption)

		switch menuOption {
		case 1:
			hangman()
			continue
		
		case 2:
			return

		default:
			fmt.Println("wrong input")
			continue
		}
	}
}

func hangman() {

	var badAttempts = 0


	var word string
	var letter string

	fmt.Println("Enter the word to play")
	fmt.Scan(&word)

	gameWord := buildhHiddenWord(word)
	originalWord := word

	for badAttempts <= 6 {
		fmt.Println("Enter the letter: ")
		fmt.Scan(&letter)

		result := compareLetter(letter, originalWord, gameWord, &badAttempts)

		gameWord = result

		fmt.Println(result)
		fmt.Printf("bad attempts: %d\n", badAttempts)
	}
}

func buildhHiddenWord(word string) []string {
	hidden := make([]string, len(word))

	for i := range word {
		hidden[i] = "_"
	}
	return hidden
}

func compareLetter(letter string, originalWord string, gameWord []string, badAttempts *int) []string {
	result := make([]string, len(gameWord))
	copy(result, gameWord)

	containsCorrectLetter := false

	for i := range originalWord {
		if letter == string(originalWord[i]) && result[i] == "_" {
			result[i] = string(originalWord[i])
			containsCorrectLetter = true
		} 	
	}

	if !containsCorrectLetter {
		*badAttempts++
	}

	return result
}

