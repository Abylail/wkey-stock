package entities

type Category struct {
	ID      int     `db:"id"`
	Code    string  `db:"code"`
	TitleRU string  `db:"title_ru"`
	TitleKZ string  `db:"title_kz"`
	Icon    *string `db:"icon"`
	Status  string  `db:"status"`
}
