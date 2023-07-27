package models

type CategoryAdminGetSingle struct {
	ID      int     `json:"id"`
	Code    string  `json:"code"`
	TitleRU string  `json:"title_ru"`
	TitleKZ string  `json:"title_kz"`
	Image   *string `json:"image"`
	Status  string  `json:"status"`

	SubCategories []SubCategoryAdminGet `json:"sub_categories"`
}

type CategoryAdminItem struct {
	ID      int     `json:"id"`
	Code    string  `json:"code"`
	TitleRU string  `json:"title_ru"`
	TitleKZ string  `json:"title_kz"`
	Image   *string `json:"image"`
	Status  string  `json:"status"`
}

// CategoryClientGet модель клиентской категории
type CategoryClientGet struct {
	Code    string  `json:"code"`
	TitleRU string  `json:"title_ru"`
	TitleKZ string  `json:"title_kz"`
	Image   *string `json:"image"`
}

type CategoryClientGetSingle struct {
	Code    string  `json:"code"`
	TitleRU string  `json:"title_ru"`
	TitleKZ string  `json:"title_kz"`
	Image   *string `json:"image"`
	Status  string  `json:"status"`
}

type CategoryAdd struct {
	TitleRU string `json:"title_ru" validate:"required"`
	TitleKZ string `json:"title_kz" validate:"required"`
}

type CategoryUpdate struct {
	TitleRU string `json:"title_ru" validate:"required"`
	TitleKZ string `json:"title_kz" validate:"required"`
}

type CategoryUpload struct {
	Image File `json:"image" validate:"required"`
}
