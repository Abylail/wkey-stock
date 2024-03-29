package entities

import "time"

type Category struct {
	ID         string `db:"id"`
	ProskladID int    `db:"prosklad_id"`
	TitleRU    string `db:"title_ru"`
	TitleKZ    string `json:"title_kz"`
	Position   int    `db:"position"`
	Count      int    `db:"count"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Deleted   bool      `db:"deleted"`
}
