package entities

import "time"

type SubCategory struct {
	ID      string `db:"id"`
	TitleRU string `db:"title_ru"`
	TitleKZ string `db:"title_kz"`
	Code    string `db:"code"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Deleted   bool      `db:"deleted"`
}
