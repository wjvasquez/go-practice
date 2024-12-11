package main

import (
	"fmt"
)

func main() {
	var year int

	fmt.Print("Enter a year: ")
	fmt.Scan(&year)

	if year <= 0 {
		fmt.Println("Error: Year has to be a positive integer.")
		return
	}

	if (year % 4 == 0 && year % 100 != 0) || (year % 400 == 0) {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}
}
