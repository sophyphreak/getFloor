package order

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
