package order

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Taxes is taxes
type Taxes struct {
	Taxes []Tax `json:"taxes"`
}

// Tax is tax
type Tax struct {
	StateAbbreviation string  `json:"stateAbbreviation"`
	StateName         string  `json:"stateName"`
	TaxRate           float64 `json:"taxRate"`
}

// Populate populates
func (t *Taxes) Populate() {
	jsonFile, err := os.Open("taxes.json")
	check(err)
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	check(err)

	json.Unmarshal(byteValue, t)
}

func (t *Taxes) getTaxRate(s string) float64 {
	for _, tax := range t.Taxes {
		if tax.StateAbbreviation == s {
			return tax.TaxRate
		}
	}
	return 0
}
