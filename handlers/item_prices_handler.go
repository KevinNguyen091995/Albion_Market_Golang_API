package handlers

import (
	"Backend/database"
	"Backend/models"
	"Backend/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMarketPriceHandler(w http.ResponseWriter, r *http.Request) {
	// Open a database connection
	db, err := database.ConnectDB()
	if err != nil {
		// Handle the error
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Execute the SQL query to get UniqueName and ItemName
	query := `SELECT "id", "last_updated", "ItemGroupTypeId", "ItemTypeId", "EnchantmentLevel", "QualityLevel", "Tier", "UnitPriceSilver" FROM market_prices_marketpricesmodel`
	rows, err := db.Query(query)

	if err != nil {
		http.Error(w, "Error querying the database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create variables to store the result
	var items []models.MarketPrice

	for rows.Next() {
		var item models.MarketPrice

		// Scan the values from the query result into variables
		scanErr := rows.Scan(&item.Id, &item.LastUpdated, &item.ItemGroupTypeId, &item.ItemTypeId, &item.EnchantmentLevel, &item.QualityLevel, &item.Tier, &item.UnitPriceSilver)
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

func GetMarketPriceDetailHandler(w http.ResponseWriter, r *http.Request) {
	db, err := database.ConnectDB()

	if err != nil {
		// Handle the error
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}

	defer db.Close()

	// Execute the SQL query to get UniqueName and ItemName
	query := `SELECT "id", "last_updated", "ItemGroupTypeId", "ItemTypeId", "EnchantmentLevel", "QualityLevel", "Tier", "UnitPriceSilver" FROM market_prices_marketpricesmodel WHERE "ItemGroupTypeId" = $1`
	params := mux.Vars(r)
	itemGroupTypeId := params["ItemGroupTypeId"]
	rows, err := db.Query(query, itemGroupTypeId)

	if err != nil {
		http.Error(w, "Error querying the database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create variables to store the result
	var items []models.MarketPrice

	for rows.Next() {
		var item models.MarketPrice

		// Scan the values from the query result into variables
		scanErr := rows.Scan(&item.Id, &item.LastUpdated, &item.ItemGroupTypeId, &item.ItemTypeId, &item.EnchantmentLevel, &item.QualityLevel, &item.Tier, &item.UnitPriceSilver)
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
