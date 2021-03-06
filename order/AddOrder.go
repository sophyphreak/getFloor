package order

// AddOrder adds an order
func (o *Orders) AddOrder(newOrder Order, t *Taxes, p *Products) {
	o.MaxOrderNumber++
	newOrder.OrderNumber = o.MaxOrderNumber
	newOrder.TaxRate = t.GetTaxRate(newOrder.State)
	cost, laborCost := p.GetProductData(newOrder.ProductType)
	newOrder.CostPerSquareFoot = cost
	newOrder.LaborCostPerSquareFoot = laborCost
	newOrder.MaterialCost = newOrder.Area * newOrder.CostPerSquareFoot
	newOrder.LaborCost = newOrder.Area * newOrder.LaborCostPerSquareFoot
	newOrder.Tax = ((newOrder.MaterialCost + newOrder.LaborCost) * (newOrder.TaxRate / 100))
	newOrder.Total = (newOrder.MaterialCost + newOrder.LaborCost + newOrder.Tax)
	o.Orders = append(o.Orders, newOrder)
	o.Commit()
}
