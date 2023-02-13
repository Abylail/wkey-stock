package entities

type BrandCreate struct {
	Title string `db:"title"`
	Image string `db:"image"`
}

type BrandUpdate struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
}

type BrandUpload struct {
	ID    int    `db:"id"`
	Image string `db:"image"`
}

type BrandGet struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
	Image string `db:"image"`
}
