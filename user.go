// user.go
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Define a simple struct to hold the data
type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var items []Item

// Get all items
func GetItems(w http.ResponseWriter, r *http.Request) {
	log.Println("GetItems endpoint hit")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// Get a single item by ID
func GetItem(w http.ResponseWriter, r *http.Request) {
	log.Println("GetItem endpoint hit")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	for _, item := range items {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.NotFound(w, r)
}

// Create a new item
func CreateItem(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateItem endpoint hit")
	w.Header().Set("Content-Type", "application/json")
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	items = append(items, item)
	json.NewEncoder(w).Encode(item)
}

// Update an item
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	log.Println("UpdateItem endpoint hit")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range items {
		if item.ID == params["id"] {
			items = append(items[:i], items[i+1:]...) // Delete old item
			var updatedItem Item
			_ = json.NewDecoder(r.Body).Decode(&updatedItem)
			updatedItem.ID = params["id"] // Keep the same ID
			items = append(items, updatedItem)
			json.NewEncoder(w).Encode(updatedItem)
			return
		}
	}
	http.NotFound(w, r)
}

// Delete an item
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteItem endpoint hit")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range items {
		if item.ID == params["id"] {
			items = append(items[:i], items[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(items)
}
