package order

import "testing"

func TestPopulateProduct(t *testing.T) {
	var p Products
	p.Populate()
	productType := p.Products[0].ProductType
	if productType != "Tile" {
		t.Errorf("Expected productType of %s but instead received %s", "Tile", productType)
	}
}

func TestGetProductData(t *testing.T) {
	var p Products
	p.Populate()
	costPerSquareFoot, laborCostPerSquareFoot := p.GetProductData("Tile")
	if costPerSquareFoot != 3.50 {
		t.Errorf("Expected costPerSquareFoot to equal %f but instead received %f", 3.50, costPerSquareFoot)
	}
	if laborCostPerSquareFoot != 4.15 {
		t.Errorf("Expected laborCostPerSquareFoot to equal %f but instead received %f", 41.5, laborCostPerSquareFoot)
	}
}
