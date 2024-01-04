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
	router.HandleFunc("/items/description", handlers.GetItemNameHandler).Methods("GET")
	router.HandleFunc("/items", handlers.GetMarketItemsHandler).Methods("GET")

    // Start the server
    http.Handle("/", router)
    http.ListenAndServe(":8080", nil)
}