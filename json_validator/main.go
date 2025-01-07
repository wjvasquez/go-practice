package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

func validJsonFormat(jsonStr string) (string, error) {
	var validStrFormat bool

	if len(jsonStr) == 0 {
		return "", errors.New("JSON string cannot be empty.")
	}

	validStrFormat = json.Valid([]byte(jsonStr))

	if !validStrFormat {
		return "", errors.New("Invalid JSON format.")
	}

	return "Valid JSON", nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	offOption := ""

	fmt.Println("JSON Validator")
	for {
		var jsonString, validStrFormat string

		fmt.Print("\nEnter a JSON string: ")
		jsonString, _ = reader.ReadString('\n')
		jsonString = strings.TrimSpace(jsonString)

		validStrFormat, err := validJsonFormat(jsonString)

		if err != nil {
			fmt.Println("Errror:", err)
			fmt.Print("\nEnter \"exit\" to finish the program or Enter to try again: ")
			fmt.Scanln(&offOption)
			if strings.ToLower(offOption) == "exit" {
				break
			}
			continue
		}

		fmt.Println(validStrFormat)
		fmt.Print("\nEnter \"exit\" to finish the program or Enter to validate another JSON string: ")
		fmt.Scanln(&offOption)
		if strings.ToLower(offOption) == "exit" {
			fmt.Println("\nThank you for using the JSON Validator!")
			break
		}
	}
}
