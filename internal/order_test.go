package order

import (
	"math"
	"testing"
)

func TestGetTotal(t *testing.T) {
	order := Order{
		Id: "1",
		Items: []Item{
			{
				Id:        "1",
				Quantity:  2,
				UnitPrice: 2.5,
			},
			{
				Id:        "2",
				Quantity:  5,
				UnitPrice: 4.99,
			},
		},
	}

	total := order.GetTotal()
	shouldBeTotal := math.Floor(float64(29.95))

	if total != shouldBeTotal {
		t.Errorf("Got %f, wanted %f", total, shouldBeTotal)
	}
}
