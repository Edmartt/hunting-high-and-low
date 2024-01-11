package main

import (
	"errors"
	"fmt"
)

type Items struct {
	price float64
}

func main() {
	item1 := Items{
		price: 109.7,
	}

	item2 := Items{
		price: 10.6,
	}

	item3 := Items{
		price: 109.9,
	}

	item4 := Items{
		price: 12.8,
	}

	items := make([]Items, 0)

	items = append(items, item4, item2, item3, item1)

	val1, val2, _ := getMinMax(items)

	fmt.Println(val1, val2)

}

func getMinMax(items []Items) (float64, float64, error) {

	if len(items) == 0 {
		return 0.0, 0.0, errors.New("emty slice")
	}

	max := items[0].price
	min := items[0].price

	for _, item := range items {
		if item.price > max {
			max = item.price
		}

		if item.price < min {
			min = item.price
		}
	}

	return max, min, nil
}
