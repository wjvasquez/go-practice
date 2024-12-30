package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func wordCount(sentence string) int {
	sentence = strings.TrimSpace(sentence)

	words := strings.Fields(sentence)

	return len(words)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nWord Counter")
	fmt.Print("Enter a sentence: ")
	sentence, _ := reader.ReadString('\n')

	wordCount := wordCount(sentence)
	fmt.Printf("The sentence contains %d words.\n", wordCount)
}
