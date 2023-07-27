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
	TitleRU       string  `json:"title_ru"`
	TitleKZ       string  `json:"title_kz"`
	DescriptionRU *string `json:"description_ru"`
	DescriptionKZ *string `json:"description_kz"`
}

type PromotionAdminUpdate struct {
	Code          string  `json:"code"`
	TitleRU       string  `json:"title_ru"`
	TitleKZ       string  `json:"title_kz"`
	DescriptionRU *string `json:"description_ru"`
	DescriptionKZ *string `json:"description_kz"`
}

type PromotionAdminUpload struct {
	Code  string `json:"code"`
	Lang  string `json:"lang"`
	Image File   `json:"image"`
}
