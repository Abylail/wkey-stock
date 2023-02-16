package entities

import "time"

type AdminProductGet struct {
	ID                int       `db:"id"`
	Title             string    `db:"title"`
	Price             int       `db:"price"`
	VendorCode        string    `db:"vendor_code"`
	Barcode           string    `db:"barcode"`
	UnitName          string    `db:"unit_name"`
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
	Position  int    `db:"position"`
	Key       string `db:"key"`
}

type ProductUpdate struct {
	ID            int    `db:"id"`
	DescriptionRU string `db:"description_ru"`
	DescriptionKZ string `db:"description_kz"`
}

type ProductUpdateImage struct {
	ProductID int    `db:"product_id"`
	Path      string `db:"path"`
	Position  int    `db:"position"`
	Key       string `db:"key"`
}

type ProductCategoryPair struct {
	ProductID       int    `db:"product_id"`
	SubCategoryID   int    `db:"sub_category_id"`
	SubCategoryName string `db:"sub_category_name"`
	SubCategoryCode string `db:"sub_category_code"`
	CategoryName    string `db:"category_name"`
	CategoryCode    string `db:"category_code"`
}
