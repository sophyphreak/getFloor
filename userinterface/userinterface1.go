package userinterface

import (
	"bufio"
	"fmt"
	"os"
)

func Menu() {
	for {

		var response string
		in := bufio.NewReader(os.Stdin)

		fmt.Println("Get Floored Program \n")
		fmt.Println("1. Display Orders")
		fmt.Println("2. Add an Order")
		fmt.Println("3. Edit an Order")
		fmt.Println("4. Remove an Order")
		fmt.Println("5. Quit\n")

		fmt.Print("Select a number 1-5: ")
		response, _ = in.ReadString('\n')

		switch response {

		case "1":
			// stuff
		case "2":
			// More stuff
		case "3":
			// Even more stuff
		case "4":

		case "5":
			fmt.Println("Thank you for coming sir")
			os.Exit(1)
		default:
			fmt.Println("Invalid response. Try again.")
			continue
		}

	}
}
