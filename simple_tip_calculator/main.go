package main

import (
	"fmt"
)

func main() {
	var bill_amount, tip_percent, tip_amount, total_amount float64

	fmt.Print("Enter total bill amount: ")
	fmt.Scan(&bill_amount)

	fmt.Print("Enter tip percentage: ")
	fmt.Scan(&tip_percent)

	tip_amount = bill_amount * (tip_percent / 100)

	total_amount = bill_amount + tip_amount

	fmt.Printf("Tip amount: $%.2f\n", tip_amount)
	fmt.Printf("Total amount: $%.2f\n", total_amount)
}
