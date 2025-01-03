package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func fileReader() (string, error) {
	var fileName string

	fmt.Print("\nEnter the file name: ")
	fmt.Scanln(&fileName)

	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", errors.New("File not found. Please check the file name and try again.")
	}

	return string(data), nil
}

func main() {
	offOption := ""
	fmt.Println("File Reader")

	for {
		content, err := fileReader()
		if err != nil {
			fmt.Println("Error:", err)

			fmt.Print("\nEnter \"exit\" to finish the program or enter to try again: ")
			fmt.Scanln(&offOption)
			if strings.ToLower(offOption) == "exit" {
				break
			}
		}

		fmt.Println(content)
	}
}
