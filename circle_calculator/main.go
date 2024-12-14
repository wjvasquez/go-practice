package main

import (
	"fmt"
	"math"
)

func main() {
	var radius, area, circumference float64

	fmt.Print("Enter the radius of the circle: ")
	fmt.Scan(&radius)

	if radius < 0 {
		fmt.Println("Error: radius must be a positive number.")
		return
	}

	area = math.Pi * math.Pow(radius, 2)
	circumference = 2 * math.Pi * radius

	fmt.Printf("Area of circle: %.2f\n", area)
	fmt.Printf("Circumference of the circle: %.2f\n", circumference)
}
