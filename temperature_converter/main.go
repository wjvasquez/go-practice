package main

import (
	"fmt"
	"strings"
)

func main() {
	var tempInput, tempConverted float64
	var givenUnit, targetUnit string
	validUnits := map[string]bool{"celsius": true, "fahrenheit": true, "kelvin": true}

	fmt.Print("Enter the temperature: ")
	fmt.Scan(&tempInput)

	fmt.Print("Enter the unit of the temperature (Celsius, Fahrenheit or Kelvin): ")
	fmt.Scan(&givenUnit)

	if !validUnits[strings.ToLower(givenUnit)] {
		fmt.Println("Error: Unsupported unit. Please use Celsius, Fahrenheit or Kelvin.")
		return
	}

	fmt.Print("Enter the target unit (Celsius, Fahrenheit or Kelvin): ")
	fmt.Scan(&targetUnit)

	if !validUnits[strings.ToLower(targetUnit)] {
		fmt.Println("Error: Unsupported unit. Please use Celsius, Fahrenheit or Kelvin.")
		return
	}

	switch strings.ToLower(givenUnit) {
	case "celsius":
		if strings.ToLower(targetUnit) == "fahrenheit" {
			tempConverted = (tempInput * 9.0 / 5.0) + 32
		} else if strings.ToLower(targetUnit) == "kelvin" {
			tempConverted = tempInput + 273.15
		} else {
			tempConverted = tempInput
		}
	case "fahrenheit":
		if strings.ToLower(targetUnit) == "celsius" {
			tempConverted = (tempInput  - 32) * (5.0 / 9.0)
		} else if strings.ToLower(targetUnit) == "kelvin" {
			tempConverted =  ( 5.0 * (tempInput - 32)) / 9.0 + 273.15
		} else {
			tempConverted = tempInput
		}
	case "kelvin":
		if strings.ToLower(targetUnit) == "celsius" {
			tempConverted = tempInput - 273.15
		} else if strings.ToLower(targetUnit) == "fahrenheit" {
			tempConverted = (9.0 / 5.0) * (tempInput - 459.67)
		} else {
			tempConverted = tempInput
		}
	}

	fmt.Printf("%.2f %s is equal to %.2f %s.\n",
			tempInput, givenUnit, tempConverted, targetUnit)

}
