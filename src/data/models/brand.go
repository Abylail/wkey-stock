package models

type Brand struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Image *string `json:"image"`
}

type BrandAdd struct {
	Title string `json:"title" validate:"required"`
	Image string `json:"image" validate:"required"`
}

type BrandUpdate struct {
	Title string `json:"title" validate:"required"`
}

type BrandUpload struct {
	Image File `json:"image" validate:"required"`
}
