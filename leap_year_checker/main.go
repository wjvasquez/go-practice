package main

import (
	"fmt"
)

func main() {
	var year int
	var isLeapYear bool

	fmt.Print("Enter a year: ")
	fmt.Scan(&year)

	if year <= 0 {
		fmt.Println("Error: Year has to be a positive integer.")
		return
	}

	if year % 4 == 0 {
		if year % 100 != 0 {
			isLeapYear = true
		} else {
			if year % 400 == 0 {
				isLeapYear = true
			} else {
				isLeapYear = false
			}
		}
	} else {
		isLeapYear = false
	}

	if isLeapYear {
		fmt.Printf("%v is a leap year.\n", year)
	} else {
		fmt.Printf("%v is not a leap year.\n", year)
	}

}
