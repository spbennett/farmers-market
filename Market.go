package main

import (
	"fmt"
)

// Market is used to track our inventory of products
// and our active discounts.
type Market struct {
	name      string
	inventory map[string]Product
	discounts map[string]Discount
}

// NewMarket is a constructor for a Market.  It will load our
// static list of active discounts/specials and populate a
// sample inventory.  These could be loaded by CSV.
func NewMarket(name string) *Market {
	fm := Market{name: name}
	fm.inventory = loadInventory()
	fm.discounts = loadDiscounts()
	return &fm
}

func (market Market) String() string {
	return fmt.Sprintf("%s", market.name)
}

// Load the expected inventory stock with the static items
// from the coding prompt.
func loadInventory() map[string]Product {
	inventory := make(map[string]Product)

	inventory["CH1"] = *NewProduct("CH1", "Chai", 3.11)
	inventory["AP1"] = *NewProduct("AP1", "Apples", 6.00)
	inventory["CF1"] = *NewProduct("CF1", "Coffee", 11.23)
	inventory["MK1"] = *NewProduct("MK1", "Milk", 4.75)
	inventory["OM1"] = *NewProduct("OM1", "Oatmeal", 3.69)

	return inventory
}

// Discount information for farmers market broken out into a generalized discount format.  The
// entries below are the discounts taken from the coding prompt, but transformed to a general format.
// qualifyingItem, qualifyingItemQuantity, discountedItem, discountedItemQuantity, discountAmountPerItem, limit
// CF1, 1, CF1, 1, 1, 0 	//1. BOGO -- Buy-One-Get-One-Free Special on Coffee. (Unlimited)
// AP1, 3, AP1, 0, 0.25, 0 	//2. APPL -- If you buy 3 or more bags of Apples, the price drops to $4.50.
// CH1, 1, MK1, 1, 1, 1		//3. CHMK -- Purchase a box of Chai and get milk free. (Limit 1)
// OM1, 1, AP1, 1, 0.50, 0	//4. APOM -- Purchase a bag of Oatmeal and get 50% off a bag of Apples
func loadDiscounts() map[string]Discount {
	discounts := make(map[string]Discount)

	discounts["BOGO"] = Discount{"BOGO", "CF1", 2, "CF1", 1, 1, 0}
	discounts["APPL"] = Discount{"APPL", "AP1", 3, "AP1", 0, 0.25, 0}
	discounts["CHMK"] = Discount{"CHMK", "CH1", 1, "MK1", 1, 1, 1}
	discounts["APOM"] = Discount{"APOM", "OM1", 1, "AP1", 1, 0.5, 0}

	return discounts
}

// printBasket pretty prints the current shopping basket
// to easily see quantities of given items.
func PrintBasket(basket []string) {
	var total int
	basketQuantity := ListToQuantityMap(basket)

	// Print contents of basket
	fmt.Println("Item\t\tQuantity")
	fmt.Println("----\t\t----")
	for key, value := range basketQuantity {
		fmt.Println(key, "\t\t", value)
		total += value
	}
	fmt.Println("-----------------------------------")
	fmt.Println("Total:\t\t", total)
}

// fillBasket validates a slice of items against our Market
// inventory.  Valid items are returned as a slice.
func FillBasket(items []string, market Market) []string {
	var result []string

	inventory := market.inventory

	for i := range items {
		var item = items[i]
		_, ok := inventory[item]
		if ok {
			result = append(result, item)
		} else {
			fmt.Println(item, "not found in inventory... skipping.")
		}
	}

	return result
}

// GetBasePrice provides a look up if the basePrice of a product
// from the Market inventory.
func (market Market) GetBasePrice(product string) float32 {
	if value, ok := market.inventory[product]; ok {
		return value.basePrice
	} else {
		return 0
	}
}

// listToQuantityMap is a helper function to get the count of each
// item in a list as a map.
func ListToQuantityMap(inputSlice []string) map[string]int {
	var quantityMap = make(map[string]int)

	for key := range inputSlice {
		var item = inputSlice[key]
		quantityMap[item]++
	}

	return quantityMap
}
