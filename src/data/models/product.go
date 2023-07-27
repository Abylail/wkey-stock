package models

import "time"

type AdminProductGet struct {
	PageCount int                `json:"page_count"`
	List      []AdminProductItem `json:"list"`
}

type AdminProductItem struct {
	ID                int                   `json:"id"`
	Title             string                `json:"title"`
	Price             int                   `json:"price"`
	VendorCode        string                `json:"vendor_code"`
	Barcode           string                `json:"barcode"`
	UnitName          string                `json:"unit_name"`
	Categories        []ProductCategoryPair `json:"categories"`
	CreatedAt         time.Time             `json:"created_at"`
	UpdatedAt         time.Time             `json:"updated_at"`
	AdditionalPercent any                   `json:"additional_percent"`
	DescriptionRU     *string               `json:"description_ru"`
	DescriptionKZ     *string               `json:"description_kz"`
	Count             int                   `json:"count"`
	Images            []string              `json:"images"`
	BrandTitle        string                `json:"brand_title"`
}

type ClientProductList struct {
	PageCount int                      `json:"page_count"`
	List      []ClientProductItemShort `json:"list"`
}

type ClientProductItemShort struct {
	ID         int      `json:"id"`
	Title      string   `json:"title"`
	Price      int      `json:"price"`
	VendorCode string   `json:"vendor_code"`
	Count      int      `json:"count"`
	Images     []string `json:"images"`
}

type ProductCategoryPair struct {
	SubCategoryCode string `json:"sub_category_code"`
	SubCategoryName string `json:"sub_category_name"`
	CategoryCode    string `json:"category_code"`
	CategoryName    string `json:"category_name"`
}

type ProductUpdate struct {
	DescriptionRU string `json:"description_ru" validate:"required"`
	DescriptionKZ string `json:"description_kz" validate:"required"`
}

type ProductUpload struct {
	Images []productImage `json:"images" validate:"required"`
}

type productImage struct {
	Position int  `json:"position" validate:"required"`
	Image    File `json:"image" validate:"required"`
}
