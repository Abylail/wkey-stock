package models

type Brand struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Code      string `json:"code"`
	Image     string `json:"image"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
