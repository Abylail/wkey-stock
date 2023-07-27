package models

type PromotionGet struct {
	ID            int     `json:"id"`
	Code          string  `json:"code"`
	TitleRU       string  `json:"title_ru"`
	TitleKZ       string  `json:"title_kz"`
	ImageRU       *string `json:"image_ru"`
	ImageKZ       *string `json:"image_kz"`
	DescriptionRU *string `json:"description_ru"`
	DescriptionKZ *string `json:"description_kz"`
}

type PromotionAdminCreate struct {
	TitleRU       string  `json:"title_ru" validate:"required"`
	TitleKZ       string  `json:"title_kz" validate:"required"`
	DescriptionRU *string `json:"description_ru" validate:"required"`
	DescriptionKZ *string `json:"description_kz" validate:"required"`
}

type PromotionAdminUpdate struct {
	Code          string  `json:"code" validate:"required"`
	TitleRU       string  `json:"title_ru" validate:"required"`
	TitleKZ       string  `json:"title_kz" validate:"required"`
	DescriptionRU *string `json:"description_ru"`
	DescriptionKZ *string `json:"description_kz"`
}

type PromotionAdminUpload struct {
	Code  string `json:"code" validate:"required"`
	Lang  string `json:"lang" validate:"required"`
	Image File   `json:"image" validate:"required"`
}
