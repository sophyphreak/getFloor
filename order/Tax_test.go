package order

import "testing"

func TestGetTaxRate(t *testing.T) {
	var taxes Taxes
	taxes.Populate()
	taxRate := taxes.GetTaxRate("OH")
	if taxRate != 6.25 {
		t.Errorf("Expected taxRate of %f but instead received %f", 6.25, taxRate)
	}
}

func TestPopulateTax(t *testing.T) {
	var taxes Taxes
	taxes.Populate()
	stateAbbrevition := taxes.Taxes[0].StateAbbreviation
	if stateAbbrevition != "OH" {
		t.Errorf("Expected stateAbbrevition of %s but instead received %s", "OH", stateAbbrevition)
	}
}
