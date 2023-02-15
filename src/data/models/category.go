package models

type CategoryAdminGet struct {
	ID      int    `json:"id"`
	Code    string `json:"code"`
	TitleRU string `json:"title_ru"`
	TitleKZ string `json:"title_kz"`
	Image   string `json:"image"`
}

type CategoryClientGet struct {
	Key   string `json:"key"`
	Title string `json:"title"`
}

type CategoryAdd struct {
	TitleRU string `json:"title_ru"`
	TitleKZ string `json:"title_kz"`
	Image   *File  `json:"image"`
}

type CategoryUpdate struct {
	TitleRU string `json:"title_ru"`
	TitleKZ string `json:"title_kz"`
}

type CategoryUpload struct {
	Image *File `json:"image"`
}
