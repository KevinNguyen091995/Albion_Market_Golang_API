package models

import (
	"time"
)

type MarketItemDescription struct {
	ID         int    `json:"-"`
	UniqueName string `json:"unique_name"`
	ItemName   string `json:"item_name"`
}

type MarketItem struct {
	ItemName        string `json:"item_name"`
	Tier            string `json:"tier"`
	ImageURL        string `json:"image_url"`
	ItemClass       string `json:"item_class"`
	ItemCategory    string `json:"item_category"`
	ItemSubCategory string `json:"item_sub_category"`
}

type MarketPrice struct {
	Id               int       `json:"-"`
	LastUpdated      time.Time `json:"last_updated"`
	ItemGroupTypeId  string    `json:"ItemGroupType"`
	ItemTypeId       string    `json:"ItemType"`
	EnchantmentLevel int       `json:"EnchantmentLevel"`
	QualityLevel     int       `json:"QualityLevel"`
	Tier             int       `json:"TierLevel"`
	UnitPriceSilver  int       `json:"UnitPriceSilver"`
}
