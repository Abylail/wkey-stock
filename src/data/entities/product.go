package entities

import "time"

type AdminProductGet struct {
	ID                int       `db:"id"`
	Title             string    `db:"title"`
	Price             int       `db:"price"`
	VendorCode        string    `db:"vendor_code"`
	Barcode           string    `db:"barcode"`
	UnitName          string    `db:"unit_name"`
	CategoryID        *int      `db:"category_id"`
	CategoryCode      *string   `db:"category_code"`
	CategoryName      *string   `db:"category_name"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
	AdditionalPercent any       `db:"additional_percent"`
	DescriptionRU     *string   `db:"description_ru"`
	DescriptionKZ     *string   `db:"description_kz"`
	Count             int       `db:"count"`
	BrandTitle        string    `db:"brand_title"`
}

type ProductImageGet struct {
	ProductID int    `db:"product_id"`
	Path      string `db:"path"`
}

type ProductUpdate struct {
	ID            int    `db:"id"`
	DescriptionRU string `db:"description_ru"`
	DescriptionKZ string `db:"description_kz"`
}
