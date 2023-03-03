package entities

type AdminPromotion struct {
	ID            int     `db:"id"`
	CODE          int     `db:"code"`
	TitleRU       *string `db:"title_ru"`
	TitleKZ       *string `db:"title_kz"`
	ImageRU       *string `db:"image_ru"`
	ImageKZ       *string `db:"image_kz"`
	DescriptionRU *string `db:"description_ru"`
	DescriptionKZ *string `db:"description_kz"`
}
