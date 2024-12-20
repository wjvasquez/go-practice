package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Create a new reader from standard input
	reader := bufio.NewReader(os.Stdin)

	var booksToAdd int
	totalInventoryValue := 0.0

	type Book struct {
		title string
		author string
		price float64
		quantity int
	}

	var books []Book

	fmt.Print("Enter the number of books to add: ")
	fmt.Scan(&booksToAdd)

	if booksToAdd <= 0 {
		fmt.Println("Error: The number of books must be a positive integer.")
		return
	}

	for i := 0; i < booksToAdd; i++ {
		book := Book{}

		// To read strings with space.
		fmt.Printf("Enter book %d title: ", i+1)
		book.title, _ = reader.ReadString('\n')
		book.title = strings.TrimSpace(book.title)

		fmt.Printf("Enter book %d author: ", i+1)
		book.author, _ = reader.ReadString('\n')
		book.author = strings.TrimSpace(book.author)

		fmt.Printf("Enter book %d price: ", i+1)
		fmt.Scan(&book.price)

		if book.price < 0 {
			fmt.Println("Error: Price must be a positive number.")
			return
		}

		fmt.Printf("Enter book %d quantity: ", i+1)
		fmt.Scan(&book.quantity)

		books = append(books, book)
		fmt.Println("")
	}

	fmt.Println("Inventory:")
	for i, book := range books {
		fmt.Printf("%d. Title: %s, Author: %s, Price: $%.2f, Quantity: %d. \n", i+1,
		book.title, book.author, book.price, book.quantity)

		totalInventoryValue = totalInventoryValue + float64(book.quantity) * book.price
	}

	fmt.Printf("Total Inventory Value: $%.2f\n", totalInventoryValue)
}
