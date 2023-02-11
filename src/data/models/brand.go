package models

type BrandCreate struct {
	Title string `json:"title"`
	Image string `json:"image"`
}

type BrandGet struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
}
