package handlers

import (
    "net/http"
	"Backend/database"
    "Backend/models"
    "Backend/utils"
)

func GetMarketItemsHandler(w http.ResponseWriter, r *http.Request) {
	// Open a database connection
	db, err := database.ConnectDB()
	if err != nil {
		// Handle the error
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Execute the SQL query to get UniqueName and ItemName
	query := "SELECT d.item_name, i.tier, i.image_url, i.item_class, i.item_category, i.item_sub_category FROM market_api_itemmodel i JOIN market_api_itemdescriptionmodel d ON i.unique_name = d.unique_name"
	rows, err := db.Query(query)

	if err != nil {
		panic("FAILED")
	}

	// Create variables to store the result
	var items []models.MarketItem

	for rows.Next() {
		var item models.MarketItem

		// Scan the values from the query result into variables
		scanErr := rows.Scan(&item.ItemName, &item.Tier, &item.ImageURL, &item.ItemClass, &item.ItemCategory, &item.ItemSubCategory)
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