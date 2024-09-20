// main.go
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Enable CORS for the API
func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func main() {
	// Initialize router
	r := mux.NewRouter()

	// Mock data (initial items)
	items = append(items, Item{ID: "1", Name: "Item One"})
	items = append(items, Item{ID: "2", Name: "Item Two"})

	// Route handlers / Endpoints
	r.HandleFunc("/items", GetItems).Methods("GET")
	r.HandleFunc("/items/{id}", GetItem).Methods("GET")
	r.HandleFunc("/items", CreateItem).Methods("POST")
	r.HandleFunc("/items/{id}", UpdateItem).Methods("PUT")
	r.HandleFunc("/items/{id}", DeleteItem).Methods("DELETE")

	// Handle CORS preflight requests
	r.Methods(http.MethodOptions).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		EnableCors(&w)
		w.WriteHeader(http.StatusOK)
	})

	// Wrap the router with CORS middleware
	http.Handle("/", r)

	// Start server on port 8000
	log.Println("Server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
