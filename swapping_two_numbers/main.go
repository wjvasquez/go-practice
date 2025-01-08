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
	fmt.Scanln(&num1)

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)

	fmt.Printf("Before swapping: a = %d, b = %d\n", num1, num2)

	swap(&num1, &num2)

	fmt.Printf("After swapping: a = %d, b = %d\n", num1, num2)
}
