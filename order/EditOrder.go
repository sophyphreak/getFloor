package order


import (
	"fmt"
)

func(o *Orders) EditOrder(orderNum int, newName, newState, newType string, newArea float64, t *Taxes, p *Products) error {
	
	foundIndex := o.findOrderIndex(orderNum)

	if foundIndex == -1 {
		return fmt.Errorf("Order number %d not found!", orderNum)
	}

	foundOrder := o.Orders[foundIndex]

	if newName != "" {
		foundOrder.CustomerName = newName
	}

	if newState != "" || newType != "" || newArea != 0{
		if newState != "" {
			foundOrder.State = newState
			foundOrder.TaxRate = t.getTaxRate(newState)
		}
		if newType != "" {
			foundOrder.ProductType = newType
		}
		if newArea != 0 {
			foundOrder.Area = newArea
		}

		cost, laborCost := p.getProductData(foundOrder.ProductType)
		foundOrder.CostPerSquareFoot = cost
		foundOrder.LaborCostPerSquareFoot = laborCost
		foundOrder.MaterialCost = foundOrder.Area * foundOrder.CostPerSquareFoot
		foundOrder.LaborCost = foundOrder.Area * foundOrder.LaborCostPerSquareFoot
		foundOrder.Tax = ((foundOrder.MaterialCost + foundOrder.LaborCost) * (foundOrder.TaxRate / 100))
		foundOrder.Total = (foundOrder.MaterialCost + foundOrder.LaborCost + foundOrder.Tax)
	}

	o.Orders[foundIndex] = foundOrder
	o.Commit()
	return nil
}