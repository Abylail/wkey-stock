package entities

type CategoryGet struct {
	ID         int    `db:"id"`
	Key        string `db:"key"`
	Title      string `db:"title"`
	ParentID   *int   `db:"parent_id"`
	Position   int    `db:"position"`
	ItemsCount *int   `db:"items_count"`
}
