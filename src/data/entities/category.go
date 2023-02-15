package entities

type CategoryGet struct {
	ID      int     `db:"id"`
	Code    string  `db:"code"`
	TitleRU string  `db:"title_ru"`
	TitleKZ string  `db:"title_kz"`
	Icon    *string `db:"icon"`
	Status  string  `db:"status"`
}

type CategoryCreate struct {
	Code     string  `db:"code"`
	TitleRU  string  `db:"title_ru"`
	TitleKZ  string  `db:"title_kz"`
	IconPath *string `db:"icon"`
}

type CategoryUpdate struct {
	Code    string `db:"code"`
	TitleRU string `db:"title_ru"`
	TitleKZ string `db:"title_kz"`
}

type CategoryUpdateImage struct {
	Code  string `db:"code"`
	Image string `db:"image"`
}
