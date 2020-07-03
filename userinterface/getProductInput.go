package userinterface

import (
	"fmt"

	"../order"
)

func getProductInput(p *order.Products) string {
start:
	prompt := "Product:"
	product := getInput(prompt)
	c, _ := p.GetProductData(product)
	if c == 0 {
		fmt.Println("Not a valid product type")
		goto start
	}
	return product
}
