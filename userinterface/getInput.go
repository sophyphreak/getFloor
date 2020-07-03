package userinterface

import "fmt"

func getInput(prompt string) string {
	var input string
	fmt.Print(prompt + " ")
	fmt.Scanln(&input)

	return input
}
