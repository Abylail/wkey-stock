package models

type SubCategoryAdminGet struct {
	ID      int     `json:"id"`
	Code    string  `json:"code"`
	TitleRU string  `json:"title_ru"`
	TitleKZ string  `json:"title_kz"`
	Image   *string `json:"image"`
	Status  string  `json:"status"`
}

type SubCategoryClientGet struct {
	Code    string  `json:"code"`
	TitleRU string  `json:"title_ru"`
	TitleKZ string  `json:"title_kz"`
	Image   *string `json:"image"`
}

type SubCategoryAdd struct {
	TitleRU string `json:"title_ru" validate:"required"`
	TitleKZ string `json:"title_kz" validate:"required"`
}

type SubCategoryUpdate struct {
	TitleRU string `json:"title_ru" validate:"required"`
	TitleKZ string `json:"title_kz" validate:"required"`
}

type SubCategoryUpload struct {
	Image *File `json:"image" validate:"required"`
}

type SubCategoryBindProductList struct {
	List []int `json:"list" validate:"required"`
}
