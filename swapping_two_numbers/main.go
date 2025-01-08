package main

import (
	"fmt"
)

func swap(a, b *int) {
	temp := *a

	*a = *b
	*b = temp
}

func main() {
	var num1, num2 int

	fmt.Print("Enter the first number: ")
	_, err := fmt.Scanln(&num1)

	if err != nil {
		fmt.Println("Error: Invalid input. Please enter a valid integer.")
		return
	}

	fmt.Print("Enter the second number: ")
	_, err = fmt.Scanln(&num2)

	if err != nil {
		fmt.Println("Error: Invalid input. Please enter a valid integer.")
		return
	}

	fmt.Printf("Before swapping: a = %d, b = %d\n", num1, num2)

	swap(&num1, &num2)

	fmt.Printf("After swapping: a = %d, b = %d\n", num1, num2)
}
