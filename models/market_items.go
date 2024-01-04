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
	ItemGroupTypeId  string    `json:"item_group"`
	ItemTypeId       string    `json:"item_type,"`
	EnchantmentLevel int       `json:"enchant_level"`
	QualityLevel     int       `json:"quality_level"`
	Tier             int       `json:"tier_level"`
	UnitPriceSilver  int       `json:"unit_price"`
}
