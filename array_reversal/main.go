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

	if err != nil {
		fmt.Println("Error: The size of the array must be an integer.")
		return
	}

	fmt.Print("Enter the elements: ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	inputDataStr = strings.Split(input, " ")

	if len(inputDataStr) != arrayLen {
		fmt.Println("Error: The number of elements is different to the Array length.")
		return
	}

	for _, element := range inputDataStr {
		intElement, _ := strconv.Atoi(element)
		originalArray = append(originalArray, intElement)
	}

	fmt.Printf("Data: %q\n", inputDataStr)
	fmt.Printf("Data: %v\n", originalArray)

	reverse(&originalArray)

	fmt.Printf("Reversal Data: %v\n", originalArray)
}
