package userinterface

import (
	"fmt"
	"os"

	"../order"
)

// Menu something
func Menu() {
	for {

		fmt.Println("Get Floored Program")
		fmt.Println("1. Display Orders")
		fmt.Println("2. Add an Order")
		fmt.Println("3. Edit an Order")
		fmt.Println("4. Remove an Order")
		fmt.Println("5. Quit")

		prompt := "Select a number 1-5:"

		var t order.Taxes
		var p order.Products
		t.Populate()
		p.Populate()
		userInput := getInput(prompt)
		switch userInput {
		case "1":
			displayOrderInput()
		case "2":
			addOrderInput(&t, &p)
		case "3":
			editOrderInput(&t, &p)
		case "4":
			deleteOrderInput()
		case "5":
			fmt.Println("Thank you for coming sir")
			os.Exit(1)
		default:
			fmt.Println("Invalid response. Try again.")
			continue
		}

	}
}
