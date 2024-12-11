package main

import (
	"fmt"
)

func main() {
	var billAmount, tipPercent, tipAmount, totalAmount float64

	fmt.Print("Enter total bill amount: ")
	fmt.Scan(&billAmount)
	if billAmount <= 0 {
		fmt.Println("Error: Bill amount must be greater than zero.")
		return
	}

	fmt.Print("Enter tip percentage: ")
	fmt.Scan(&tipPercent)
	if tipPercent < 0 {
		fmt.Println("Error: Tip percentage cannot be negative.")
		return
	}

	tipAmount = billAmount * (tipPercent / 100)

	totalAmount = billAmount + tipAmount

	fmt.Printf("Tip amount: $%.2f\n", tipAmount)
	fmt.Printf("Total amount: $%.2f\n", totalAmount)
}
