package order

import (
	"encoding/json"
	"io/ioutil"
)

// Commit something
func (o *Orders) Commit() error {
	f, err := json.MarshalIndent(o, "", "\t")
	check(err)

	err = ioutil.WriteFile("Orders_"+o.Date+".json", f, 0644)
	return err
}
