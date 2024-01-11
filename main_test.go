package main

import "testing"

type testTable struct {
	array         []Items
	expectedMax   float64
	expectedMin   float64
	expectedError error
}

func TestGetMinMax(t *testing.T) {

	tables := []testTable{
		{
			array:         []Items{{price: 7.9}, {price: 12.4}, {price: 150.8}, {price: 64.9}},
			expectedMax:   150.8,
			expectedMin:   7.9,
			expectedError: nil,
		},
		{
			array:         []Items{{price: 190.8}, {price: 5.8}, {price: 1.8}, {price: 55.8}},
			expectedMax:   190.8,
			expectedMin:   1.8,
			expectedError: nil,
		},
	}

	for _, table := range tables {
		max, min, err := getMinMax(table.array)

		if err != table.expectedError {
			t.Errorf("wrong error: expected %v and got %v", table.expectedError, err)
		}

		if max != table.expectedMax || min != table.expectedMin {
			t.Errorf("wrong results: expected values: %v, %v and got %v, %v", table.expectedMax, table.expectedMin, max, min)
		}
	}

}
