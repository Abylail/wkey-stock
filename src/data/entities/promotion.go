package entities

import (
	"time"
)

type Promotion struct {
	ID         string `db:"id"`
	Link       string `db:"link"`
	LinkTextRU string `db:"link_text_ru"`
	LinkTextKZ string `db:"link_text_kz"`
	Image      string `db:"image"`
	Position   int    `db:"position"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Deleted   bool      `db:"deleted"`
}
