package main

import (
	"fmt"
)

func main() {
	type Product struct {
		Name     string
		Price    float64
		Quantity int
	}

	var menuOption int
	var finalBill float64
	isMenuActive := true
	var shoppingCart []Product

	fmt.Printf("\nWelcome to the Simple Shopping Cart!\n")

	for {

		fmt.Printf("\n1. Add Product\n")
		fmt.Printf("2. View Cart\n")
		fmt.Printf("3. Checkout\n")

		fmt.Print("\nChoose an option: ")
		fmt.Scanln(&menuOption)

		switch menuOption {
		case 1:
			product := Product{}

			fmt.Print("\nEnter Product name: ")
			fmt.Scanln(&product.Name)
			fmt.Print("Enter product price: ")
			fmt.Scanln(&product.Price)

			if product.Price <= 0 {
				fmt.Println("Error: Price must be a positive value.")
				return
			}

			fmt.Print("Enter product quantity: ")
			fmt.Scanln(&product.Quantity)

			if product.Quantity <= 0 {
				fmt.Println("Error: Quantity must be greater than zero.")
				return
			}

			shoppingCart = append(shoppingCart, product)

			fmt.Printf("\nProduct added to the cart!\n")

		case 2:
			fmt.Printf("\nCart:\n")
			for i, product := range shoppingCart {
				fmt.Printf("%d. %s - $%.2f X %d = $%.2f\n", i+1, product.Name,
					product.Price, product.Quantity, float64(product.Quantity)*product.Price)
			}

		case 3:
			fmt.Printf("\nFinal Bill:\n")
			for i, product := range shoppingCart {
				fmt.Printf("%d. %s - $%.2f X %d = $%.2f\n", i+1, product.Name,
					product.Price, product.Quantity, float64(product.Quantity)*product.Price)

				finalBill += float64(product.Quantity) * product.Price
			}

			fmt.Printf("Total: $%.2f\n", finalBill)
			isMenuActive = false

		default:
			fmt.Println("Invalid option, try again!")
		}

		if isMenuActive == false {
			fmt.Printf("\nThank you for shopping with us!\n")
			break
		}
	}
}
