package main

import (
	"testing"
)

// Use test data from code puzzle prompt.
var testData = []struct {
	inputItems []string
	expected   float32
}{
	{[]string{"CH1", "AP1", "CF1", "MK1"}, 20.34},
	{[]string{"MK1", "AP1"}, 10.75},
	{[]string{"CF1", "CF1"}, 11.23},
	{[]string{"AP1", "AP1", "CH1", "AP1"}, 16.61},
	{[]string{"CF1", "CF1", "CF1", "CF1"}, 22.46},
	{[]string{"CF1", "CF1", "CF1"}, 22.46},
}

func TestBaskets(t *testing.T) {
	fm := NewMarket("The Farmers Market")

	// Fill the shopping basket with items from testBaskets.
	for _, value := range testData {
		items := value.inputItems
		basket := fillBasket(items, *fm)
		register := checkout(basket, *fm)
		actual := register.GetTotal()
		if actual != value.expected {
			t.Errorf("Basket(%s): expected %f, actual %f", value.inputItems, value.expected, actual)
		}
	}
}
