package main

import (
	"fmt"
)

func main() {
	var menuOption, productQuantity int
	var productName string
	var productPrice, finalBill float64
	isMenuActive := true
	shoppingCart := make(map[string]float64)

	fmt.Printf("Welcome to the Simple Shopping Cart!\n\n")

	for {

		fmt.Printf("1. Add Product\n")
		fmt.Printf("2. View Cart\n")
		fmt.Printf("3. Checkout\n\n")

		fmt.Print("Choose an option: ")
		fmt.Scan(&menuOption)

		switch menuOption {
		case 1:
			fmt.Print("\nEnter Product name: ")
			fmt.Scanln(&productName)
			fmt.Print("Enter product price: ")
			fmt.Scanln(&productPrice)
			fmt.Print("Enter product quantity: ")
			fmt.Scanln(&productQuantity)

			shoppingCart[productName] = productPrice

			fmt.Printf("\nProduct added to the cart!\n\n")

		case 2:
			i := 1

			fmt.Printf("\nCart:\n")
			for product, price := range shoppingCart {
				fmt.Printf("%d. %s - $%.2f X %d = $%.2f\n\n", i, product, price,
					productQuantity, float64(productQuantity) * price)
				i += 1
			}

		case 3:
			i := 1

			fmt.Printf("\nFinal Bill:\n")
			for product, price := range shoppingCart {
				fmt.Printf("%d. %s - $%.2f X %d = $%.2f\n", i, product, price,
					productQuantity, float64(productQuantity) * price)

				finalBill += float64(productQuantity) * price

				i += 1
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
