package order

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

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

// Populate populates
func (p *Products) Populate() {
	jsonFile, err := os.Open("products.json")
	check(err)
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	check(err)

	json.Unmarshal(byteValue, p)
}

func (p *Products) getProductData(productType string) (float64, float64) {
	for _, v := range p.Products {
		if v.ProductType == productType {
			return v.CostPerSquareFoot, v.LaborCostPerSquareFoot
		}
	}
	return 0, 0
}
