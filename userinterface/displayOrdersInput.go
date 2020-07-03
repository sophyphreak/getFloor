package userinterface

import (
	"fmt"

	"../order"
)

func displayOrderInput() {
	var o order.Orders
	date := GetDate()
	o.Populate(date)
	err := o.Display()
	if err != nil {
		fmt.Println(err)
	}
}
