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
		lines := processDiscount(discounts[key], basketQuantity, inventory)
		register.lines = append(register.lines, lines...)
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

// processDiscount takes basket items that qualify for the discount and moves them to registry entries.
// Remaining items are left in basket for other discounts.
func processDiscount(discount Discount, basketQuantity map[string]int, inventory map[string]Product) []Line {
	var lines []Line
	var entry Line
	var discountAppliedCount = 0

	// Unpack current discount.
	var qualifyingItem = discount.qualifyingItem
	var qualifyingItemQuantity = discount.qualifyingItemQuantity
	var discountedItem = discount.discountedItem
	var discountedItemQuantity = discount.discountedItemQuantity
	var discountPerDiscountedItem = discount.discountPerDiscountedItem
	var limit = discount.limit

	// Stop applying this discount when we reach our discount limit.
	for discountAppliedCount = 0; discountAppliedCount <= limit; discountAppliedCount++ {

		// Check if enough qualifying items are found.
		if basketQuantity[qualifyingItem] >= qualifyingItemQuantity && basketQuantity[discountedItem] >= discountedItemQuantity {
			var discountPrice float32
			var discountItemBasePrice = inventory[discountedItem].basePrice
			var qualifyingItemBasePrice = inventory[qualifyingItem].basePrice

			// When discountedItemQuantity is zero, all discountedItems receive the discount.
			if discountedItemQuantity == 0 {

				// Calculate discount to apply (Avoid divide by zero error).
				discountPrice = discountPerDiscountedItem * discountItemBasePrice

				// Consume all of the discounted items.
				for i := basketQuantity[qualifyingItem]; i > 0; i-- {
					basketQuantity[qualifyingItem]--
					entry = Line{discountedItem, discountItemBasePrice, discount.id, discountPrice * -1}
					lines = append(lines, entry)
				}
			} else {

				// Calculate discount to apply (Avoid divide by zero error).
				discountPrice = discountPerDiscountedItem * float32(discountedItemQuantity) * discountItemBasePrice

				// Consume discounted items.
				basketQuantity[discountedItem] = basketQuantity[discountedItem] - discountedItemQuantity
				entry = Line{discountedItem, discountItemBasePrice, discount.id, discountPrice * -1}
				lines = append(lines, entry)

				// Consume the qualifying items.
				basketQuantity[qualifyingItem]--
				entry = Line{qualifyingItem, qualifyingItemBasePrice, "", 0}
				lines = append(lines, entry)
			}

			// Loop unlimited specials.
			if limit == 0 {
				discountAppliedCount--
			}
		}
	}
	return lines
}