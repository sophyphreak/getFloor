package userinterface

import (
	"fmt"

	"../order"
)

func addOrderInput(t *order.Taxes, p *order.Products) {
	fmt.Println("Please enter the date you want the order for")
	date := GetDate()
	name := getNameInput()
	product := getProductInput(p)
	state := getStateInput(t)
	area := getAreaInput()
	var o order.Orders
	o.Populate(date)
	newOrder := order.Order{0, name, state, 0, product, area, 0, 0, 0, 0, 0, 0}
	o.AddOrder(newOrder, t, p)
	fmt.Println("Order successfully added.")
}
