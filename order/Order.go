package order

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Orders is the orders
type Orders struct {
	Orders         []Order `json:"orders"`
	Date           string
	MaxOrderNumber int
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

// Populate populates
func (o *Orders) Populate(date string) {
	o.Date = date
	filename := "Orders_" + date + ".json"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		os.Create(filename)
	}
	jsonFile, err := os.Open(filename)
	check(err)
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	check(err)

	json.Unmarshal(byteValue, o)

	max := 0
	for _, v := range o.Orders {
		if v.OrderNumber > max {
			max = v.OrderNumber
		}
	}
	o.MaxOrderNumber = max
}
