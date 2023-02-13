package entities

type SubCategoryGet struct {
	ID      int    `db:"id"`
	Code    string `db:"code"`
	TitleRU string `db:"title_ru"`
	TitleKZ string `db:"title_kz"`
	Icon    string `db:"icon"`
}

type SubCategoryCreate struct {
	Code     string `db:"code"`
	TitleRU  string `db:"title_ru"`
	TitleKZ  string `db:"title_kz"`
	IconPath string `db:"icon"`
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
