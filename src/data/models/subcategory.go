package models

type SubCategory struct {
	ID      string `json:"id"`
	TitleRU string `json:"title_ru"`
	TitleKZ string `json:"title_kz"`
	Code    string `json:"code"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
