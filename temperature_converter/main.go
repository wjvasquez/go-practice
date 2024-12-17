package main

import (
	"fmt"
	"strings"
)

func main() {
	var tempInput, tempConverted float64
	var givenUnit, targetUnit string

	fmt.Print("Enter the temperature: ")
	fmt.Scan(&tempInput)

	fmt.Print("Enter the unit of the temperature (Celsius, Fahrenheit or Kelvin): ")
	fmt.Scan(&givenUnit)

	if strings.ToLower(givenUnit) != "celsius" && strings.ToLower(givenUnit) != "fahrenheit" &&
	strings.ToLower(givenUnit) != "kelvin" {
		fmt.Println(givenUnit)
		fmt.Println("Error: Unsupported unit. Please use Celsius, Fahrenheit or Kelvin.")
		return
	}

	fmt.Print("Enter the target unit (Celcius, Fahrenheit or Kelvin): ")
	fmt.Scan(&targetUnit)

	if strings.ToLower(targetUnit) != "celsius" && strings.ToLower(targetUnit) != "fahrenheit" &&
	strings.ToLower(targetUnit) != "kelvin" {
		fmt.Println("Error: Unsupported unit. Please use Celsius, Fahrenheit or Kelvin.")
		return
	}

	switch givenUnit {
	case "celsius":
		if strings.ToLower(targetUnit) == "fahrenheit" {
			tempConverted = (tempInput * 9 / 5) + 32
		} else if strings.ToLower(targetUnit) == "kelvin" {
			tempConverted = tempInput + 273.15
		} else {
			tempConverted = tempInput
		}
	case "fahrenheit":
		if strings.ToLower(targetUnit) == "celsius" {
			tempConverted = (tempInput  - 32) * (5 / 9)
		} else if strings.ToLower(targetUnit) == "kelvin" {
			tempConverted =  ( 5 * (tempInput - 32)) / 9 + 273.15
			fmt.Println("Hello!")
		} else {
			tempConverted = tempInput
		}
	case "kelvin":
		if strings.ToLower(targetUnit) == "celsius" {
			tempConverted = tempInput - 273.15
		} else if strings.ToLower(targetUnit) == "fahrenheit" {
			tempConverted = (9 / 5) * (tempInput - 459.67)
		} else {
			tempConverted = tempInput
		}
	}

	fmt.Printf("%.2f %s is equal to %.2f %s.\n",
			tempInput, givenUnit, tempConverted, targetUnit)

}
