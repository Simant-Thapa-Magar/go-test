package order

import "math"

type Item struct {
	Id        string
	Quantity  int
	UnitPrice float64
}

type Order struct {
	Id    string
	Items []Item
}

func (o Order) GetTotal() float64 {
	var total float64

	for _, item := range o.Items {
		total += item.UnitPrice * float64(item.Quantity)
	}

	return math.Floor(total)
}