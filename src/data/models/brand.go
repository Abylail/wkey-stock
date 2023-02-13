package models

type BrandAdd struct {
	Title string `json:"title"`
	Image string `json:"image"`
}

type BrandUpdate struct {
	Title string `json:"title"`
}

type BrandUpload struct {
	Image *File `json:"image"`
}

type BrandGet struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
}
