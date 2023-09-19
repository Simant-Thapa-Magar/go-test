package order

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
	shouldBeTotal := 29.95

	if total != shouldBeTotal {
		t.Errorf("Got %f, wanted %f", total, shouldBeTotal)
	}
}

func TestGetTotalUsingTestify(t *testing.T) {
	orders := []Order{
		{
			Id: "1",
			Items: []Item{
				{
					Id:        "1",
					Quantity:  10,
					UnitPrice: 2.49,
				},
				{
					Id:        "2",
					Quantity:  5,
					UnitPrice: 1.99,
				},
			},
		},
		{
			Id: "2",
			Items: []Item{
				{
					Id:        "3",
					Quantity:  1,
					UnitPrice: 49.95,
				},
				{
					Id:        "4",
					Quantity:  7,
					UnitPrice: 5,
				},
			},
		},
	}

	shouldBeTotal := []float64{34.85, 84.95}

	for index, order := range orders {
		total := order.GetTotal()
		fmt.Println("what's total ", total)
		assert.Equal(t, total, shouldBeTotal[index], "Total should be equal")
	}
}
