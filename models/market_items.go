package models

type MarketItemDescription struct {
	ID int `json:"-"`
    UniqueName string `json:"unique_name"`
	ItemName string `json:"item_name"`
}

type MarketItem struct {
	ItemName string `json:"item_name"`
	Tier string `json:"tier"`
	ImageURL string `json:"image_url"`
	ItemClass string `json:"item_class"`
	ItemCategory string `json:"item_category"`
	ItemSubCategory string `json:"item_sub_category"`
}