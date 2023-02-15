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
	Key   string `json:"key"`
	Title string `json:"title"`
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
