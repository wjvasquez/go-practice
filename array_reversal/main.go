package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reverse(arr *[]int) {
	arrLength := len(*arr)
	for i := 0; i < arrLength/2; i++ {
		temp := (*arr)[i]
		(*arr)[i] = (*arr)[arrLength-1-i]
		(*arr)[arrLength-1-i] = temp
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var arrayLen int
	var input string
	var inputDataStr []string
	var originalArray []int

	fmt.Print("Enter the size of the array: ")
	_, err := fmt.Scanln(&arrayLen)

	if arrayLen <= 0 {
		fmt.Println("Error: Array size must be a positive integer.")
		return
	}

	if err != nil {
		fmt.Println("Error: The size of the array must be an integer.")
		return
	}

	fmt.Print("Enter the elements: ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	inputDataStr = strings.Split(input, " ")

	if len(inputDataStr) != arrayLen {
		fmt.Printf("Error: Expected %d elements, but got %d.\n", arrayLen, len(inputDataStr))
		return
	}

	for _, element := range inputDataStr {
		intElement, err := strconv.Atoi(element)
		if err != nil {
			fmt.Printf("Error: '%s' is not a valid integer.\n", element)
			return
		}
		originalArray = append(originalArray, intElement)
	}

	fmt.Printf("\nOriginal Array: %v\n", originalArray)
	fmt.Println("\nReversing array...")
	reverse(&originalArray)
	fmt.Printf("\nReversed Array: %v\n", originalArray)
}
