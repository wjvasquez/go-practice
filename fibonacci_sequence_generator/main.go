package main

import (
	"fmt"
	"strconv"
	"strings"
)

func fibonacci(n int) []int {
	fibonacciSeq := []int{0, 1}

	if n == 1 {
		return []int{0}
	}

	for i := 2; i < n; i++ {
		fibonacciSeq = append(fibonacciSeq, fibonacciSeq[i-1]+fibonacciSeq[i-2])
	}

	return fibonacciSeq
}

func fibonacciStr(fibonacciSeq []int) []string {
	var fibonacciSeqStr []string

	for _, term := range fibonacciSeq {
		fibonacciSeqStr = append(fibonacciSeqStr, strconv.Itoa(term))
	}

	return fibonacciSeqStr
}

func main() {
	var terms int

	fmt.Printf("\nFibonacci Sequence Generator\n")
	fmt.Print("Enter the number of terms: ")
	fmt.Scanln(&terms)

	if terms <= 0 {
		fmt.Println("\nError: The input must be a positive integer.")
		return
	}

	fmt.Printf("\nFibonacci sequence: %s\n", strings.Join(fibonacciStr(fibonacci(terms)), ", "))
}
