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
		fmt.Println("3-> Delete a note")
		fmt.Println("4-> Exit")
		fmt.Printf("Enter your choice: ")
		fmt.Scan(&switchChoice)
		fmt.Printf("Your choice %d\n", switchChoice)

		switch switchChoice {

		case 1:
			writeFile()
			continue

		case 2:
			readFile()
			continue

		case 3:
			deleteFile()
			continue

		case 4:
			fmt.Println("Loop ended")
			return

		default:
			fmt.Println("Error.")
		}
	}	
}


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
