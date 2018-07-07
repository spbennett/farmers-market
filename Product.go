package main

// Product information for farmers market.
type Product struct {
	productCode string
	name        string
	basePrice   float32
}

// NewProduct creates a new product where
// productCode is a string representing a product id
// name is a string representing a friendly name for the product
// basePrice is a float32 representing the price of a product
func NewProduct(productCode string, name string, basePrice float32) *Product {
	product := Product{productCode, name, basePrice}
	return &product
}
