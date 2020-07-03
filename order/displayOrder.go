package order

import "fmt"

func (o *Orders) Display() error {

	for _, ord := range o.Orders {
		fmt.Println(ord.OrderNumber, o.Date)
		fmt.Println(ord.CustomerName)
		fmt.Println(ord.State)
		fmt.Println("Product:", ord.ProductType)
		fmt.Printf("Materials: %.2f\n", ord.MaterialCost)
		fmt.Printf("Labor: %.2f\n", ord.LaborCost)
		fmt.Printf("Tax: %.2f\n", ord.Tax)
		fmt.Printf("Total: %.2f\n\n", ord.Total)
	}
	if len(o.Orders) == 0 {
		return fmt.Errorf("That date does not have any orders")
	}

	return nil
}
