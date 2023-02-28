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
	Code  string  `json:"code"`
	Title string  `json:"title"`
	Image *string `json:"image"`
}

type SubCategoryAdd struct {
	TitleRU string `json:"title_ru"`
	TitleKZ string `json:"title_kz"`
}

type SubCategoryUpdate struct {
	TitleRU string `json:"title_ru"`
	TitleKZ string `json:"title_kz"`
}

type SubCategoryUpload struct {
	Image *File `json:"image"`
}

type SubCategoryBindProductList struct {
	List []int `json:"list"`
}
