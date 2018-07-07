package main

type Discount struct {
	id                        string
	qualifyingItem            string
	qualifyingItemQuantity    int
	discountedItem            string
	discountedItemQuantity    int
	discountPerDiscountedItem float32
	limit                     int
}

