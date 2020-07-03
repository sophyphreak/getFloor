package userinterface

import (
	"fmt"
	"strings"

	"../order"
)

func editOrderInput(t *order.Taxes, p *order.Products) {
	var name, state, product string
	var area float64
	date := GetDate()
	var o order.Orders
	o.Populate(date)
	orderNumber := getOrderNumber(&o)
	fmt.Println("You can change the name, the state, the product, or the area.")
update:
	prompt := "What would you like to update?"
	userInput := getInput(prompt)
	userInput = strings.ToLower(userInput)
	switch userInput {
	case "name", "nam", "na", "n":
		name = getNameInput()
		o.EditOrder(orderNumber, name, state, product, area, t, p)
		goto update
	case "state", "stat", "sta", "st", "s":
		state = getStateInput(t)
		o.EditOrder(orderNumber, name, state, product, area, t, p)
		goto update
	case "product", "prod", "pro", "pr", "p":
		product = getProductInput(p)
		o.EditOrder(orderNumber, name, state, product, area, t, p)
		goto update
	case "area", "are", "ar", "a":
		area = getAreaInput()
		o.EditOrder(orderNumber, name, state, product, area, t, p)
		goto update
	case "quit", "qui", "qu", "q":
		break
	default:
		goto update
	}
}
