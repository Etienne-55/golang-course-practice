package projects

import (
	"os"
	"fmt"
	"bufio"
) 


func Notepad() {

for {
	var switchChoice int
		fmt.Println("----Notepad----")
		fmt.Println("1-> Enter a new note")
		fmt.Println("2-> Search a note ")
		fmt.Println("3-> Edit a note")
		fmt.Println("6-> Exit")
		fmt.Printf("Enter your choice: ")
		fmt.Scan(&switchChoice)
		fmt.Printf("Your choice %d\n", switchChoice)

		switch switchChoice {

		case 1:
			writeFile()

		case 2:
			readFile()

		case 6:
			fmt.Println("Loop ended")
			return

		default:
			fmt.Println("Error.")
		}
	}	
}


func writeFile() error {
	// var fileName string
	// var note string
	// fmt.Println("Enter the note name: ")
	// fmt.Scan(&fileName)
	//
	// formatedFileName := fmt.Sprintf("\"%s.txt\"", fileName)
	//
	// fmt.Printf("Enter the note: ")
	// scanner := bufio.NewScanner(os.Stdin)	
	// if scanner.Scan() {
	// 	note = scanner.Text()
	// 	os.WriteFile(formatedFileName, []byte(note), 0644)
	// }

	
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


// func readFile() {
// var fileName string
//
// 	fmt.Printf("Enter the name of the note you want to read: ")
// 	fmt.Scan(&fileName)
// 	formatedFileName := fmt.Sprintf("\"%s.txt\"", fileName)
//
// 	noteData, err := os.ReadFile(formatedFileName)
// 	if err != nil {
// 		fmt.Println("Error while reading the file, file not found.")
// 	}
//
// 	fileContent := string(noteData)
// 	fmt.Println(fileContent)
// }


func readFile() error {
	var fileName string
	fmt.Print("Enter the name of the file you want to read: ")
	if _, err := fmt.Scanln(&fileName); err != nil {
		return fmt.Errorf("failed to read filename: %w", err)
	}

	formattedFileName := fileName + ".txt"
	noteData, err := os.ReadFile(formattedFileName)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	fileContent := string(noteData)
	fmt.Println(fileContent)

	return nil
}


func deleteFile() error {
	var fileName string
	fmt.Printf("Enter the file you want to delete: ")
	if _, err := fmt.Scanln(&fileName); err != nil {
		return fmt.Errorf("failed to read filename: %w", err)
	}

	formattedFileName := fileName + ".txt"
	err := os.Remove(formattedFileName)
	if err != nil {
		return fmt.Errorf("failed to delete, file not found")
	}

	fmt.Println("File deleted successfully!")
	return nil
}
