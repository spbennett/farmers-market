package main

// Product information for farmers market.
type Product struct {
	productCode string
	name        string
	basePrice	float32
}

// Create a new product
// productCode string A product id
// name string A friendly name for the product
// basePrice float32 The price of a product
func NewProduct (productCode string, name string, basePrice float32) *Product {
	product := Product{productCode, name, basePrice}
	return &product
}