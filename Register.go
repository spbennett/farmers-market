package main

import (
	"fmt"
	"strings"
)

// Register stores our itemized bill with discounts/specials.
type Register struct {
	lines []Line
}

// Line represents one possible entry in our itemized bill.
type Line struct {
	item          string
	itemBasePrice float32
	discount      string
	discountPrice float32
}

// Lines returns the lines field of the given Register.
func Lines(register Register) []Line {
	return register.lines
}

// GetTotal returns the current total for a register.
func (register Register) GetTotal() float32 {
	var total float32 = 0
	var lines = register.lines

	for key := range lines {
		var itemBasePrice = lines[key].itemBasePrice
		var discountPrice = lines[key].discountPrice
		total += itemBasePrice + discountPrice
	}

	return total
}

// Pretty print the register contents to match the formatting in the code
// prompt.
func (register Register) String() string {
	var total float32 = 0
	var lines = register.lines

	builder := strings.Builder{}

	// Print contents of basket
	builder.WriteString("Item\t\tPrice\n")
	builder.WriteString("----\t\t----\n")
	for key := range lines {
		var item = lines[key].item
		var itemBasePrice = lines[key].itemBasePrice
		var discount = lines[key].discount
		var discountPrice = lines[key].discountPrice

		builder.WriteString(fmt.Sprintf("%s\t\t%6.2f\n", item, itemBasePrice))
		if discountPrice == 0 {
			builder.WriteString(fmt.Sprintf("\t%s\t\n", discount))
		} else {
			builder.WriteString(fmt.Sprintf("\t%s\t%6.2f\n", discount, discountPrice))
		}
		total += itemBasePrice + discountPrice
	}
	builder.WriteString("-----------------------------------\n")
	builder.WriteString(fmt.Sprintf("Total:\t\t%6.2f\n", register.GetTotal()))
	return builder.String()
}

// checkout transforms a grocery basket into a receipt by checking the basket
// for each active discount/special.  As specials are applied, items are moved
// from basket to register.  Any remaining items after discount are tallied up
// into the receipt.
func checkout(basket []string, market Market) Register {
	discounts := market.discounts
	inventory := market.inventory
	var register Register

	basketQuantity := listToQuantityMap(basket)

	// Iterate through each discount and apply to our basket.
	for key := range discounts {
		var discount = discounts[key]
		var discountAppliedCount = 0

		// Unpack current discount.
		var qualifyingItem = discount.qualifyingItem
		var qualifyingItemQuantity = discount.qualifyingItemQuantity
		var discountedItem = discount.discountedItem
		var discountedItemQuantity = discount.discountedItemQuantity
		var discountPerDiscountedItem = discount.discountPerDiscountedItem

		// Check if enough qualifying items are found.
		if basketQuantity[qualifyingItem] >= qualifyingItemQuantity && basketQuantity[discountedItem] > 0 {
			var qualifyingItemBasePrice = inventory[qualifyingItem].basePrice
			var discountItemBasePrice = inventory[discountedItem].basePrice
			var discountItemQuantity = discount.discountedItemQuantity

			// When discountedItemQuantity is zero, there are no limits to applying the discount.
			if discountItemQuantity == 0 {
				var discountPrice = discount.discountPerDiscountedItem * discountItemBasePrice
				for i := basketQuantity[qualifyingItem]; i > 0; i-- {
					var entry = Line{qualifyingItem, qualifyingItemBasePrice, discount.id, discountPrice * -1}
					register.lines = append(register.lines, entry)
					basketQuantity[qualifyingItem]--
				}
			} else {
				var discountPrice = discountPerDiscountedItem * float32(discountItemQuantity) * discountItemBasePrice

				var entry = Line{qualifyingItem, qualifyingItemBasePrice, discount.id, discountPrice * -1}
				register.lines = append(register.lines, entry)

				// Remove discounted items from being counted twice.
				basketQuantity[qualifyingItem] = basketQuantity[qualifyingItem] - discountedItemQuantity
			}

			// Apply a discount.
			discountAppliedCount++
		}

		// Stop applying this discount when we reach our discount limit.  Limit of 0 is unlimited.
		if discountAppliedCount >= discount.limit && discount.limit != 0 {
			continue
		}
	}

	// Zero out remaining basket items with no discount.
	for key := range basketQuantity {
		for i := 0; i < basketQuantity[key]; i++ {
			var item = key
			var basePrice = inventory[item].basePrice

			var entry = Line{item, basePrice, "", 0}
			register.lines = append(register.lines, entry)
		}
	}

	return register
}
