package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var myScore, computerScore int
	var validOptions = [3]string{"rock", "paper", "scissors"}
	var results = [3]string{"win", "lose", "tied"}

	for {
		var myChoice, computerChoice, gameDecision string
		var result string

		fmt.Println("\nLet's Play Rock, Paper, Scissors!")
		fmt.Print("Enter your choice (rock, paper, scissors): ")
		fmt.Scanln(&myChoice)

		computerChoice = validOptions[rand.Intn(3)]

		switch myChoice {
		case "rock":
			switch computerChoice {
			case "rock":
				result = results[2]

			case "paper":
				result = results[1]

			case "scissors":
				result = results[0]
			}
		case "paper":
			switch computerChoice {
			case "rock":
				result = results[0]

			case "paper":
				result = results[2]

			case "scissors":
				result = results[1]
			}
		case "scissors":
			switch computerChoice {
			case "rock":
				result = results[1]

			case "paper":
				result = results[0]

			case "scissors":
				result = results[2]
			}
		}

		switch result {
		case "win":
			myScore++
		case "lose":
			computerScore++
		}

		fmt.Printf("Computer chose %s.\n", computerChoice)
		fmt.Printf("You %s this round!\n", result)
		fmt.Printf("\nScore: You %d - %d Computer\n", myScore, computerScore)

		fmt.Printf("Do you want to play again? (yes/no): ")
		fmt.Scanln(&gameDecision)

		if gameDecision == "no" {
			fmt.Printf("Thanks for playing!\n")
			break
		}
	}

}
