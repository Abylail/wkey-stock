package entities

type Brand struct {
	ID         int     `db:"id"`
	Title      string  `db:"title"`
	Image      *string `db:"image"`
	ProskladID int     `db:"prosklad_id"`
}
