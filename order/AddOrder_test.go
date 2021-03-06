package order

import "testing"

func TestAddOrder(t *testing.T) {
	var o Orders
	var taxes Taxes
	var p Products
	taxes.Populate()
	p.Populate()

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

	order := Order{0, customerName, state, 0, productType, area, 0, 0, 0, 0, 0, 0}

	o.AddOrder(order, &taxes, &p)

	foundOrder := o.Orders[0]
	if foundOrder.OrderNumber != orderNumber {
		t.Errorf("Got OrderNumber %d but expected %d", foundOrder.OrderNumber, orderNumber)
	}
	if foundOrder.CustomerName != customerName {
		t.Errorf("Got CustomerName %s but expected %s", foundOrder.CustomerName, customerName)
	}
	if foundOrder.State != state {
		t.Errorf("Got State %s but expected %s", foundOrder.State, state)
	}
	if foundOrder.TaxRate != taxRate {
		t.Errorf("Got TaxRate %f but expected %f", foundOrder.TaxRate, taxRate)
	}
	if foundOrder.ProductType != productType {
		t.Errorf("Got ProductType %s but expected %s", foundOrder.ProductType, productType)
	}
	if foundOrder.Area != area {
		t.Errorf("Got Area %f but expected %f", foundOrder.Area, area)
	}
	if foundOrder.CostPerSquareFoot != costPerSquareFoot {
		t.Errorf("Got CostPerSquareFoot %f but expected %f", foundOrder.CostPerSquareFoot, costPerSquareFoot)
	}
	if foundOrder.LaborCostPerSquareFoot != laborCostPerSquareFoot {
		t.Errorf("Got LaborCostPerSquareFoot %f but expected %f", foundOrder.LaborCostPerSquareFoot, laborCostPerSquareFoot)
	}
	if foundOrder.MaterialCost != materialCost {
		t.Errorf("Got MaterialCost %f but expected %f", foundOrder.MaterialCost, materialCost)
	}
	if foundOrder.LaborCost != laborCost {
		t.Errorf("Got LaborCost %f but expected %f", foundOrder.LaborCost, laborCost)
	}
	if foundOrder.Tax != tax {
		t.Errorf("Got Tax %f but expected %f", foundOrder.Tax, tax)
	}
	if foundOrder.Total != total {
		t.Errorf("Got Total %f but expected %f", foundOrder.Total, total)
	}
}
