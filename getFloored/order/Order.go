package order

// Orders is the orders
type Orders struct {
	Orders []Order `json:"orders"`
}

// Order is an order
type Order struct {
	OrderNumber            int     `json:"orderNumber"`
	CustomerName           string  `json:"customerName"`
	State                  string  `json:"state"`
	TaxRate                float64 `json:"taxRate"`
	ProductType            string  `json:"productType"`
	Area                   float64 `json:"area"`
	CostPerSquareFoot      float64 `json:"costPerSquareFoot"`
	LaborCostPerSquareFoot float64 `json:"laborCostPerSquareFoot"`
	MaterialCost           float64 `json:"materialCost"`
	LaborCost              float64 `json:"laborCost"`
	Tax                    float64 `json:"tax"`
	Total                  float64 `json:"total"`
}
