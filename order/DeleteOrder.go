package order


import (
	"fmt"
)



func(o *Orders) DeleteOrder(ordNum int) error{

	foundIndx := o.findOrderIndex(ordNum) 
	if foundIndx != -1 {
		lastIndx := len(o.Orders) - 1
		if lastIndx == 0 { //if only one order
			o.Orders = o.Orders[:0]
		} else { //if more than one order
			o.Orders[foundIndx], o.Orders[lastIndx] = o.Orders[lastIndx], o.Orders[foundIndx]
			o.Orders = o.Orders[:lastIndx]
		}
		o.Commit()
		return nil
	}
	return fmt.Errorf("Order %d not found!", ordNum)
}