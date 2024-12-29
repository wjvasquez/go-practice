package main

import (
	"fmt"
	"strconv"
	"strings"
)

func factorial(number int) int {
	factorialNum := 1

	for i := 1; i <= number; i++ {
		factorialNum = i * factorialNum
	}

	return factorialNum
}

func main() {
	fmt.Printf("Factorial Calcultor\n")

	for {
		var input string
		fmt.Print("\nEnter a number (Enter \"exit\" to finish the app): ")
		fmt.Scanln(&input)

		if strings.ToLower(input) == "exit" {
			fmt.Println("\nThank you for use the Factorial Calculator App!")
			return
		}

		var number int
		number, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("\nError: Invalid input. Please enter a valid integer.")
			continue
		}

		if number < 0 {
			fmt.Println("\nError: Factorial is not defined for negative numbers.")
			continue
		}

		fmt.Printf("Factorial of %d is %d.\n", number, factorial(number))
	}
}
