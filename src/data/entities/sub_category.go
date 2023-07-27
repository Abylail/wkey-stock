package entities

type SubCategory struct {
	ID       int     `db:"id"`
	Code     string  `db:"code"`
	TitleRU  string  `db:"title_ru"`
	TitleKZ  string  `db:"title_kz"`
	Icon     *string `db:"icon"`
	Status   string  `db:"status"`
	ParentID int     `db:"parent_id"`
}
