package entities

type SubCategoryGet struct {
	ID      int     `db:"id"`
	Code    string  `db:"code"`
	TitleRU string  `db:"title_ru"`
	TitleKZ string  `db:"title_kz"`
	Icon    *string `db:"icon"`
	Status  string  `db:"status"`
}

type SubCategoryCreate struct {
	Code     string `db:"code"`
	TitleRU  string `db:"title_ru"`
	TitleKZ  string `db:"title_kz"`
	ParentID int    `db:"parent_id"`
}

type SubCategoryUpdate struct {
	Code    string `db:"code"`
	TitleRU string `db:"title_ru"`
	TitleKZ string `db:"title_kz"`
}

type SubCategoryUpdateImage struct {
	Code  string `db:"code"`
	Image string `db:"image"`
}
