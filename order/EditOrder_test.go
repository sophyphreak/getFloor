package order

import (
	"testing"
)
func TestEditOrder(t *testing.T) {
	var o Orders
	o.Populate("07032020")

	var taxes Taxes
	var p Products
	taxes.Populate()
	p.Populate()


	err := o.EditOrder(1, "", "", "", 500, &taxes, &p)

	if err != nil {
		t.Errorf("Expected nil error but received %v!", err)
	}
}