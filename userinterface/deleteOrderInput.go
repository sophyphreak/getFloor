package userinterface

import (
	"fmt"
	"strings"

	"../order"
)

func deleteOrderInput() {
	date := GetDate()
	var o order.Orders
	o.Populate(date)
	orderNumber := getOrderNumber(&o)
	index := o.FindOrderIndex(orderNumber)
	if index == -1 {
		fmt.Println(fmt.Errorf("Order number does not exist in this date"))
		return
	}
	prompt := "Are you sure you want to delete this order?"
	ans := getInput(prompt)
	ans = strings.ToLower(ans)
	switch ans {
	case "yes", "ye", "y":
		err := o.DeleteOrder(orderNumber)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Order successfully deleted.")
		}
	default:
		fmt.Println("Order was not deleted")
	}
}
