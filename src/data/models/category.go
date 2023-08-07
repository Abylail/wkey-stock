package models

type Category struct {
	ID         string `json:"id"`
	ProskladID int    `json:"prosklad_id"`
	TitleRU    string `json:"title_ru"`
	TitleKZ    string `json:"title_kz"`
	Position   int    `json:"position"`
	Count      int    `json:"count"`
}

type CategoryProsklad struct {
	ID       int    `json:"id" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Position int    `json:"position" validate:"required"`
	Count    int    `json:"count"`
}

type CategoryUpdateCount struct {
	Count int `json:"count" validate:"required"`
}
