package models

type Promotion struct {
	ID         string `json:"id"`
	Link       string `json:"link"`
	LinkTextRU string `json:"link_text_ru"`
	LinkTextKZ string `json:"link_text_kz"`
	Image      string `json:"image"`
	Position   int    `json:"position"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
