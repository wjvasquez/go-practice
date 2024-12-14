package main

import (
	"fmt"
	"strings"
)

func main() {
	const exchangeRate = 4000
	var moneyAmount, convertedAmount float64
	var sourceCurrency, targetCurrency string

	fmt.Print("Enter the amount to convert: ")
	fmt.Scan(&moneyAmount)

	if moneyAmount <= 0 {
		fmt.Println("Error: Please enter an amount greater than 0.")
		return
	}

	fmt.Print("Choose source currency (COP): ")
	fmt.Scan(&sourceCurrency)

	fmt.Print("Choose target currency (USD): ")
	fmt.Scan(&targetCurrency)

	if strings.ToUpper(sourceCurrency) != "COP" || strings.ToUpper(targetCurrency) != "USD" {
		fmt.Println("Error: Sopported conversion is only from COP to USD.")
		return
	}

	convertedAmount = moneyAmount / exchangeRate

	fmt.Printf("%.2f %s is equal to %.2f %s\n",
			moneyAmount, strings.ToUpper(sourceCurrency),
			convertedAmount, strings.ToUpper(targetCurrency))
}
