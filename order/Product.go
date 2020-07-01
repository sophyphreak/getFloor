package order

// Products is Products
type Products struct {
	Products []Product `json:"products"`
}

// Product is Product
type Product struct {
	ProductType            string  `json:"productType"`
	CostPerSquareFoot      float64 `json:"costPerSquareFoot"`
	LaborCostPerSquareFoot float64 `json:"laborCostPerSquareFoot"`
}
