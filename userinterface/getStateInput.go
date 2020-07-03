package userinterface

import (
	"fmt"
	"strings"

	"../order"
)

func getStateInput(t *order.Taxes) string {
start:
	prompt := "State Abbreviation:"
	state := getInput(prompt)
	state = strings.ToUpper(state)
	if len(state) != 2 {
		fmt.Println("Please provide two letters")
		goto start
	}
	tax := t.GetTaxRate(state)
	if tax == 0 {
		fmt.Println("No data found for that state")
		goto start
	}
	return state
}
