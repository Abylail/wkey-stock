package entities

type CategoryGet struct {
	ID      int    `db:"id"`
	Code    string `db:"code"`
	TitleRU string `db:"title_ru"`
	TitleKZ string `db:"title_kz"`
	Icon    string `db:"icon"`
}

type CategoryCreate struct {
	Code     string `json:"code"`
	TitleRU  string `db:"title_ru"`
	TitleKZ  string `db:"title_kz"`
	IconPath string `db:"icon"`
}

type CategoryUpdate struct {
	ID      int    `db:"id"`
	Code    string `db:"code"`
	TitleRU string `db:"title_ru"`
	TitleKZ string `db:"title_kz"`
	Icon    string `db:"icon"`
}
