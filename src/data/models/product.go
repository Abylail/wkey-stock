package models

type Product struct {
	ID            string  `json:"id"`
	ProskladID    int     `json:"prosklad_id"`
	Title         string  `json:"title"`
	DescriptionRU *string `json:"description_ru"`
	DescriptionKZ *string `json:"description_kz"`
	Count         int     `json:"count"`
}

type ProductProsklad struct {
	// common
	ID               int    `json:"id" validate:"required"`
	Title            string `json:"title" validate:"required"`
	Barcode          string `json:"barcode" validate:"required"`
	CompanyID        int    `json:"company_id" validate:"required"`
	ItemCategoryName string `json:"item_category_name"`

	// money
	SellingPrice          float32 `json:"selling_price" validate:"required"`
	OldSellingPrice       float32 `json:"old_selling_price" validate:"required"`
	PreviousPurchasePrice float32 `json:"previous_purchase_price"`
	AdditionalPercent     float32 `json:"additional_percent"`

	// flags
	HasInventory bool `json:"has_inventory"`
	IsVirtual    bool `json:"is_virtual"`
	Marked       bool `json:"marked"`
	IsQuick      bool `json:"is_quick"`

	// unit
	UnitID   int    `json:"unit_id" validate:"required"`
	UnitName string `json:"unit_name" validate:"required"`
	UnitType int    `json:"unit_type" validate:"required"`

	// vendor
	VendorID   int    `json:"vendor_id"`
	VendorCode string `json:"vendor_code" validate:"required"`
}

type ProductUpdateDescription struct {
	DescriptionRU string `json:"description_ru" validate:"required"`
	DescriptionKZ string `json:"description_kz" validate:"required"`
}

type ProductUpdateCount struct {
	Count int `json:"count" validate:"required"`
}
