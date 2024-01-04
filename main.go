// main.go
package main

import (
	"net/http"

	"Backend/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Define API routes
	//Items GET
	router.HandleFunc("/items/all", handlers.GetMarketItemsHandler).Methods("GET")
	router.HandleFunc("/items/detail/description", handlers.GetItemNameHandler).Methods("GET")
	router.HandleFunc("/items/detail/categories", handlers.GetItemCategoryListHandler).Methods("GET")
	router.HandleFunc("/items/{ItemCategory}", handlers.GetItemCategoryHandler).Methods("GET")

	//Prices GET
	router.HandleFunc("/prices/all", handlers.GetMarketPriceHandler).Methods("GET")
	router.HandleFunc("/prices/detail/{ItemGroupTypeId}", handlers.GetMarketPriceDetailHandler).Methods("GET")

	// Start the server
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
