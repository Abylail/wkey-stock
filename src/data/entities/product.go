package entities

import "time"

type Product struct {
	// common
	ID               string `db:"id"`
	ProskladID       int    `db:"prosklad_id"`
	Title            string `db:"title"`
	Barcode          string `db:"barcode"`
	CompanyID        int    `db:"company_id"`
	ItemCategoryName string `db:"item_category_name"`

	// price
	SellingPrice          float32 `db:"selling_price"`
	OldSellingPrice       float32 `db:"old_selling_price"`
	PreviousPurchasePrice float32 `db:"previous_purchase_price"`
	AdditionalPercent     float32 `db:"additional_percent"`

	// flags
	HasInventory bool `db:"has_inventory"`
	IsVirtual    bool `db:"is_virtual"`
	Marked       bool `db:"marked"`
	IsQuick      bool `db:"is_quick"`

	// unit
	UnitID   int    `db:"unit_id"`
	UnitName string `db:"unit_name"`
	UnitType int    `db:"unit_type"`

	// brand (vendor)
	BrandID int `db:"brand_id"`

	// custom
	DescriptionRU *string   `db:"description_ru"`
	DescriptionKZ *string   `db:"description_kz"`
	Count         int       `db:"count"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}
