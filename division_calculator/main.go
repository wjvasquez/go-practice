package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getNumber() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("Error reading input: %v", err)
	}

	input = strings.TrimSpace(input)

	number, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("Invalid input. Please enter a valid number.")
	}

	return number, nil
}

func div(numerator int, denominator int) (float64, error) {
	if denominator == 0 {
		return 0, errors.New("Division by zero is not allowed.")
	}

	result := float64(numerator) / float64(denominator)

	return result, nil
}

func main() {
	offOption := ""
	fmt.Println("\nDivision Calcultor")

	for {
		var numerator, denominator int
		var result float64

		fmt.Print("\nEnter numerator: ")
		numerator, err := getNumber()

		if err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}

		fmt.Print("Enter denominator: ")
		denominator, err = getNumber()

		if err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}

		result, err = div(numerator, denominator)

		if err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}

		fmt.Printf("Result: %.2f.\n", result)

		fmt.Printf("\nEnter \"exit\" to finish the program: ")
		fmt.Scanln(&offOption)

		if strings.ToLower(offOption) == "exit" {
			fmt.Println("Thank you for use the Division Calculator.")
			return
		}
	}
}
