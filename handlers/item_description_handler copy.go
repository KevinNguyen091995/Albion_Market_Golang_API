package handlers

import (
    "net/http"
	"Backend/database"
    "Backend/models"
    "Backend/utils"
)


func GetItemNameHandler(w http.ResponseWriter, r *http.Request) {
	// Open a database connection
	db, err := database.ConnectDB()
	if err != nil {
		// Handle the error
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Execute the SQL query to get UniqueName and ItemName
	query := "SELECT unique_name, item_name FROM market_api_itemdescriptionmodel"
	rows, err := db.Query(query)

	if err != nil {
		panic("FAILED")
	}

	// Create variables to store the result
	var items []models.MarketItemDescription

	for rows.Next() {
		var item models.MarketItemDescription

		// Scan the values from the query result into variables
		scanErr := rows.Scan(&item.UniqueName, &item.ItemName)
		if scanErr != nil {
			// handle error
			utils.RespondError(w, http.StatusInternalServerError, "Error querying the database")
			return
		}

		// Use append to add the item to the items slice
		items = append(items, item)
	}


	// Respond with the items in JSON format
	utils.RespondJSON(w, http.StatusOK, items)
}