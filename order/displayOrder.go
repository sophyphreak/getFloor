package order

import "fmt"

func (o *Orders) Display() error {

	for _, ord := range o.Orders {
		fmt.Println(ord.OrderNumber, o.Date)
		fmt.Println(ord.CustomerName)
		fmt.Println(ord.State)
		fmt.Println("Product:", ord.ProductType)
		fmt.Println("Materials:", ord.MaterialCost)
		fmt.Println("Labor:", ord.LaborCost)
		fmt.Println("Tax:", ord.Tax)
		fmt.Println("Total", ord.Total)
	}
	if len(o.Orders) == 0 {
		return fmt.Errorf("That date does not have any orders")
	}

	return nil
}
