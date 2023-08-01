package entities

import "time"

type Category struct {
	ID         string    `db:"id"`
	ProskladID int       `db:"prosklad_id"`
	Title      string    `db:"title"`
	Position   int       `db:"position"`
	Count      int       `db:"count"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	Deleted    bool      `db:"deleted"`
}
