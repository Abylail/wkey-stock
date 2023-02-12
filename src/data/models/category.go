package models

type CategoryAdminGet struct {
	ID      int    `json:"id"`
	Code    string `json:"code"`
	TitleRU string `json:"title_ru"`
	TitleKZ string `json:"title_kz"`
	Icon    string `json:"icon"`
}

type CategoryClientGet struct {
	Key   string `json:"key"`
	Title string `json:"title"`
}

type CategoryAdd struct {
	TitleRU string `json:"title_ru"`
	TitleKZ string `json:"title_kz"`
	Icon    string `json:"icon"`
}

type CategoryUpdate struct {
	TitleRU string `json:"title_ru"`
	TitleKZ string `json:"title_kz"`
	Icon    string `json:"icon"`
}
