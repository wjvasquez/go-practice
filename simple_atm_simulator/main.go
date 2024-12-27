package main

import (
	"fmt"
)

func main() {
	var balance, depositAmount, withdrawalAmount float64
	var menuOption int
	isMenuActive := true

	fmt.Printf("\nWelcome to the simple ATM!\n")

	for {
		fmt.Printf("\n1. Check Balance\n")
		fmt.Printf("2. Deposit Money\n")
		fmt.Printf("3. Withdraw Money\n")
		fmt.Printf("4. Exit\n")

		fmt.Print("Choose an option: ")
		fmt.Scanln(&menuOption)

		switch menuOption {
		case 1:
			fmt.Printf("\nYour current balance is $%.2f.\n", balance)

		case 2:
			fmt.Printf("\nEnter deposit amount: ")
			fmt.Scanln(&depositAmount)

			if depositAmount <= 0 {
				fmt.Printf("Error: The money to deposit must be a positive number. Try again!\n")
				continue
			}

			balance += depositAmount

			fmt.Printf("Deposit successful. New balance: $%.2f.\n", balance)

		case 3:
			fmt.Printf("\nEnter withdrawal amount: ")
			fmt.Scanln(&withdrawalAmount)

			if withdrawalAmount > balance {
				fmt.Printf("\nError: Insuficient balance.\n")
				continue
			}

			balance -= withdrawalAmount

			fmt.Printf("Withdrawal successful. New balance: $%.2f.\n", balance)
		case 4:
			fmt.Printf("\nThank you for use our service!\n")
			isMenuActive = false
		default:
			fmt.Printf("\nInvalid option, enter a number between 1 and 4.\n")
		}

		if !isMenuActive {
			break
		}
	}
}
