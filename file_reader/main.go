package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func readFileContent(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", errors.New("File not found. Please check the file name and try again.")
	}

	content := strings.TrimSpace(string(data))
	if len(content) == 0 {
		return "", errors.New("The file is empty.")
	}

	return content, nil
}

func main() {
	offOption := ""
	fmt.Println("File Reader")

	for {
		var fileName string
		fmt.Print("Enter the file name: ")
		fmt.Scanln(&fileName)

		content, err := readFileContent(fileName)
		if err != nil {
			fmt.Println("Error:", err)
			fmt.Print("\nEnter \"exit\" to finish the program or press Enter to try again: ")
			fmt.Scanln(&offOption)
			if strings.ToLower(offOption) == "exit" {
				break
			}
			continue
		}

		fmt.Print("\nFile Content:\n\n")
		fmt.Println(content)

		fmt.Print("\nEnter \"exit\" to finish the program or press Enter to read another file: ")
		fmt.Scanln(&offOption)
		if strings.ToLower(offOption) == "exit" {
			fmt.Println("\nThank you for using the File Reader!")
			break
		}
	}
}
