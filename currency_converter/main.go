package main

import (
	"fmt"
	"strings"
)

func main() {
	const exchangeRate = 4000
	var moneyAmout, moneyConverted float64
	var sourceCurrency, targetCurrency string

	fmt.Print("Enter the amount to convert: ")
	fmt.Scan(&moneyAmout)

	if moneyAmout < 0 {
		fmt.Println("Error: The amount to convert must be a positive number.")
		return
	}

	fmt.Print("Choose source currency (COP): ")
	fmt.Scan(&sourceCurrency)

	fmt.Print("Choose target currency (USD): ")
	fmt.Scan(&targetCurrency)

	moneyConverted = moneyAmout / exchangeRate

	fmt.Printf("%.2f %s is equal to %.2f %s\n", moneyAmout, strings.ToUpper(sourceCurrency), moneyConverted, strings.ToUpper(targetCurrency))
}
