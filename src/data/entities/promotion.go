package entities

type AdminPromotion struct {
	ID            int     `db:"id"`
	CODE          string  `db:"code"`
	TitleRU       string  `db:"title_ru"`
	TitleKZ       string  `db:"title_kz"`
	ImageRU       *string `db:"image_ru"`
	ImageKZ       *string `db:"image_kz"`
	DescriptionRU *string `db:"description_ru"`
	DescriptionKZ *string `db:"description_kz"`
}

type AdminPromotionCreate struct {
	Code          string  `db:"code"`
	TitleRU       string  `db:"title_ru"`
	TitleKZ       string  `db:"title_kz"`
	DescriptionRU *string `db:"description_ru"`
	DescriptionKZ *string `db:"description_kz"`
}

type AdminPromotionUpdate struct {
	Code          string  `db:"code"`
	TitleRU       string  `db:"title_ru"`
	TitleKZ       string  `db:"title_kz"`
	DescriptionRU *string `db:"description_ru"`
	DescriptionKZ *string `db:"description_kz"`
}
