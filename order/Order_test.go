package order

import (
	"os"
	"testing"
)

// Orders_MMDDYYYY.json

func TestPopulateOrders(t *testing.T) {
	var o Orders
	date := "11112020"
	o.Populate(date)

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
	tax := 61.88
	total := 1051.88

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

	date = "12122020"
	o.Populate(date)
	filename := "Orders_" + date + ".json"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Errorf("File %s was not created", filename)
	}
	os.Remove(filename)
}

func TestFindOrderIndex(t *testing.T) {
	var o Orders
	o.Populate("07032020")

	indx := o.FindOrderIndex(1)

	if indx == -1 {
		t.Errorf("Got indx -1, but expected 1")
	}
}
