package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var items []string

// Message stores IDs from incoming HTTP messages
// {"id": "CF1"}
type Message struct {
	ID string `json:"id"`
}

// Serve two main paths: /add and /checkout.
func main() {
	http.HandleFunc("/add", basketHandler)
	http.HandleFunc("/checkout", checkoutHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// checkoutHandler takes an existing list of items,
// validates them into a shopping basket, then
// processes discounts and displays output.
func checkoutHandler(w http.ResponseWriter, r *http.Request) {
	fm := NewMarket("The Farmers Market")

	// Move added items to the shopping basket.
	basket := fillBasket(items, *fm)
	items = nil

	register := checkout(basket, *fm)
	fmt.Fprintf(w, register.String())
}

// basketHandler takes ids from json and stores them
// in our list of items.  These ids will be thrown away
// later if they are not products in Market's inventory.
func basketHandler(w http.ResponseWriter, r *http.Request) {
	reqError := "Example request: {\"id\": \"CF1\"}"
	var msg Message
	if r.Body == nil {
		http.Error(w, reqError, 400)
		return
	}
	// Decode incoming json.
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, reqError, 400)
		return
	}
	fmt.Println(fmt.Sprintf("%s received", msg.ID))

	// Append new item to list.
	items = append(items, msg.ID)
}
