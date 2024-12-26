package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var maxAttempts int
	var initialInput string
	randomNumber := rand.Intn(100) + 1
	attempts := 0

	fmt.Printf("\nI'm thinking of a number between 1 and 100. Can you guess it?\n")
	fmt.Print("How many attempts do you want to get? ")
	fmt.Scanln(&maxAttempts)

	if maxAttempts <= 0 {
		fmt.Println("Please enter a positive number of attempts.")
		return
	}

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

		if guessNumber < 1 || guessNumber > 100 {
			fmt.Println("Please guess a number between 1 and 100.")
		}

		if guessNumber == randomNumber {
			fmt.Printf("\nCorrect! You've guessed the number in %d attempts.\n", attempts)
			break
		} else if guessNumber < randomNumber {
			fmt.Printf("Too low! Try again. You have %d attempts left.\n", maxAttempts-attempts)
		} else {
			fmt.Printf("Too high! Try again. Yuo have %d attempts left.\n", maxAttempts-attempts)
		}

		if attempts == maxAttempts {
			fmt.Printf("\nYou've reached the maximum number of attempts.\n")
			fmt.Printf("The number was %d.\n", randomNumber)
			break
		}
	}
}
