package order

import (
	"testing"
)

func TestDeleteOrder(t *testing.T) {
	var o Orders
	o.Populate("07032020")
	// defer os.Remove("Orders_07032020.json")

	orderNumber := 1
	customerName := "Wise"
	state := "OH"
	productType := "Wood"
	area := 100.00

	taxRate := 6.25
	laborCostPerSquareFoot := 4.75
	costPerSquareFoot := 5.15
	materialCost := area * costPerSquareFoot
	laborCost := area * laborCostPerSquareFoot
	tax := ((materialCost + laborCost) * (taxRate / 100))
	total := (materialCost + laborCost + tax)

	order := Order{orderNumber, customerName, state, taxRate, productType, area, costPerSquareFoot, laborCostPerSquareFoot, materialCost, laborCost, tax, total}
	o.Orders = append(o.Orders, order)

	err := o.Commit()
	check(err)

	err = o.DeleteOrder(1)

	if err != nil{
		t.Errorf("OrderNumber 1 was not found")
	}
}

func TestDeleteOrder1(t *testing.T) {
	var o Orders
	o.Populate("07032020")
	// defer os.Remove("Orders_07032020.json")

	orderNumber := 1
	customerName := "Wise"
	state := "OH"
	productType := "Wood"
	area := 100.00

	taxRate := 6.25
	laborCostPerSquareFoot := 4.75
	costPerSquareFoot := 5.15
	materialCost := area * costPerSquareFoot
	laborCost := area * laborCostPerSquareFoot
	tax := ((materialCost + laborCost) * (taxRate / 100))
	total := (materialCost + laborCost + tax)

	order := Order{orderNumber, customerName, state, taxRate, productType, area, costPerSquareFoot, laborCostPerSquareFoot, materialCost, laborCost, tax, total}
	o.Orders = append(o.Orders, order)

	err := o.Commit()
	check(err)

	err = o.DeleteOrder(2)

	if err == nil{
		t.Errorf("Expected error: \"Order 2 not found\" but received %v error", err)
	}
}

