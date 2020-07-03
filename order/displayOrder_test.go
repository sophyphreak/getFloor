package order

import (
	"testing"
)

func TestDisplay(t *testing.T) {
	var o Orders
	date := "12122020"
	o.Populate(date)
	err := o.Display()

	if err == nil {
		t.Errorf("Got error %v but expected That date does not have any orders", err)
	}
}
