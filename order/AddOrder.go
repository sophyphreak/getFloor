package order

// AddOrder adds an order
func (o *Orders) AddOrder(newOrder Order) {
	o.Orders = append(o.Orders, newOrder)
	// o.Commit()
}
