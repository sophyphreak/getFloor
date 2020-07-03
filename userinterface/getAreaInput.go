package userinterface

import (
	"fmt"
	"strconv"
)

func getAreaInput() float64 {
area:
	prompt := "Area:"
	userInput := getInput(prompt)
	a, err := strconv.ParseFloat(userInput, 64)
	if err != nil {
		fmt.Println("Input is not a float")
		goto area
	}
	if a < 100 {
		fmt.Println("area must be 100 or above")
		goto area
	}
	return a
}
