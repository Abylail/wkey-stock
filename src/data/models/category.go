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
	Code  string  `json:"code"`
	Title string  `json:"title"`
	Image *string `json:"image"`
}

type CategoryAdd struct {
	TitleRU string `json:"title_ru"`
	TitleKZ string `json:"title_kz"`
}

type CategoryUpdate struct {
	TitleRU string `json:"title_ru"`
	TitleKZ string `json:"title_kz"`
}

type CategoryUpload struct {
	Image File `json:"image"`
}
