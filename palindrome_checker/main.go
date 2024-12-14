package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var originalString, stringLetters string

	fmt.Print("Enter a string: ")
	originalString, _ = reader.ReadString('\n')
	originalString = strings.TrimSpace(originalString)

	if originalString == "" {
		fmt.Printf("Empty input is considered a palindrome.\n")
		return
	}

	stringLetters = strings.ToLower(strings.ReplaceAll(originalString, " ",""))

	// remove punctuation or symbols from the letters.
	for _, r := range stringLetters {
		// if there is a symbol is removed from the string to compare.
		if !unicode.IsLetter(r) {
			stringLetters = strings.ReplaceAll(stringLetters, string(r), "")
		}
	}

	for i := 0; i < len(stringLetters); i++ {
		// compare the letters from start to end and reverse.
		if stringLetters[i] != stringLetters[len(stringLetters) - i - 1] {
			fmt.Printf("%q is not a palindrome\n", stringLetters)
			return
		}
	}

	fmt.Printf("%q is a palindrome\n", originalString)
}
