package entities

import "time"

type Brand struct {
	ID    string `db:"id"`
	Title string `db:"title"`
	Code  string `db:"code"`
	Image string `db:"image"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Deleted   bool      `db:"deleted"`
}
