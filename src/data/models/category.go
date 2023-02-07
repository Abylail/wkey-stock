package models

type CategoryAdminGet struct {
	ID         int    `json:"id"`
	Key        string `json:"key"`
	Title      string `json:"title"`
	ParentID   *int   `json:"parent_id"`
	Position   int    `json:"position"`
	ItemsCount *int   `json:"items_count"`
}

type CategoryClientGet struct {
	Key   string `json:"key"`
	Title string `json:"title"`
}
