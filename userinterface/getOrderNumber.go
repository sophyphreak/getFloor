package userinterface

import (
	"fmt"
	"strconv"

	"../order"
)

func getOrderNumber(o *order.Orders) int {
start:
	prompt := "Order number:"
	userInput := getInput(prompt)
	orderNumber, err := strconv.Atoi(userInput)
	if err != nil || orderNumber < 0 {
		fmt.Println(fmt.Errorf("Please provide whole number"))
		goto start
	}
	index := o.FindOrderIndex(orderNumber)
	if index == -1 {
		fmt.Println("Order not found")
		goto start
	}
	return orderNumber
}
