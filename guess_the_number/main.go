package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func main() {
	var maxAttempts int
	var initialInput string
	randomNumber := rand.Intn(100) + 1
	attempts := 0

	fmt.Printf("\nI'm thinking of a number between 1 and 100. Can you guess it?\n")
	fmt.Print("How many attempts do you want to get? ")
	fmt.Scanln(&maxAttempts)

	for {
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&initialInput)

		guessNumber, err := strconv.Atoi(initialInput)

		if err != nil {
			if strings.ToLower(initialInput) == "exit" {
				fmt.Printf("\nThanks for playing! Goodbye.\n")
				return
			}

			fmt.Printf("\nInvalid input. Please enter a number between 1 and 100.\n")
			continue
		}

		attempts++

		if guessNumber == randomNumber {
			fmt.Printf("\nCorrect! You've guessed the number in %d attempts.\n", attempts)
			break
		} else if guessNumber < randomNumber {
			fmt.Println("Too low! Try again.")
		} else {
			fmt.Println("Too high! Try again.")
		}

		if attempts == maxAttempts {
			fmt.Printf("\nYou've reached the maximum number of attemps.\n")
			fmt.Printf("The number was %d.\n", randomNumber)
			break
		}
	}
}
