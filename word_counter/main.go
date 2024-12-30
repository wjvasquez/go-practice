package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func wordCount(sentence string) int {
	sentence = strings.TrimSpace(sentence)

	if len(sentence) == 0 {
		return 0
	}

	word := ""
	words := 1

	for _, chacarter := range sentence {
		word += string(chacarter)

		if unicode.IsSpace(chacarter) {
			words++
			word = ""
		}
	}

	return words
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var sentence string

	fmt.Println("\nWord Counter")
	fmt.Print("Enter a sentence: ")
	sentence, _ = reader.ReadString('\n')

	fmt.Printf("The sentence contains %d words.\n", wordCount(sentence))
}
