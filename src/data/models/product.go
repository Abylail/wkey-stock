package models

import "time"

type AdminProductGet struct {
	ID                int       `json:"id"`
	Title             string    `json:"title"`
	Price             int       `json:"price"`
	VendorCode        string    `json:"vendor_code"`
	Barcode           string    `json:"barcode"`
	UnitName          string    `json:"unit_name"`
	CategoryID        int       `json:"category_id"`
	CategoryName      string    `json:"category_name"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	AdditionalPercent any       `json:"additional_percent"`
	DescriptionRU     *string   `json:"description_ru"`
	DescriptionKZ     *string   `json:"description_kz"`
	Count             int       `json:"count"`
}
