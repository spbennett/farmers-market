package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)

var items []string

type Message struct {
	Id   string  `json:"id"`
}

func main() {
	http.HandleFunc("/add", basketHandler)
	http.HandleFunc("/checkout", checkoutHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func checkoutHandler(w http.ResponseWriter, r *http.Request) {
	fm := NewMarket("The Farmers Market")

	// Move added items to the shopping basket.
	basket := fillBasket(items, *fm)
	items = nil

	register := checkout(basket, *fm)
	fmt.Fprintf(w, register.String())
}

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
	fmt.Println(fmt.Sprintf("%s received", msg.Id))

	// Append new item to list.
	items = append(items, msg.Id)
}





