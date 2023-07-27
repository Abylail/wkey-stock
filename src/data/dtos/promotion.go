package dtos

import (
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
	"wkey-stock/src/enums/languages"
)

type Promotion struct {
	id            int
	code          string
	titleRU       string
	titleKZ       string
	imageRU       *string
	imageKZ       *string
	descriptionRU *string
	descriptionKZ *string
}

func NewPromotion(entity *entities.Promotion) *Promotion {
	return &Promotion{
		id:            entity.ID,
		code:          entity.Code,
		titleRU:       entity.TitleRU,
		titleKZ:       entity.TitleKZ,
		imageRU:       entity.ImageRU,
		imageKZ:       entity.ImageKZ,
		descriptionRU: entity.DescriptionRU,
		descriptionKZ: entity.DescriptionKZ,
	}
}

func (promotion Promotion) ID() int {
	return promotion.id
}

func (promotion Promotion) Code() string {
	return promotion.code
}

func (promotion Promotion) Title(language string) string {
	if language == languages.KZ {
		return promotion.titleKZ
	}

	return promotion.titleRU
}

func (promotion Promotion) Image(language string) *string {
	if language == languages.KZ {
		return promotion.imageKZ
	}

	return promotion.imageRU
}

func (promotion Promotion) Description(language string) *string {
	if language == languages.KZ {
		return promotion.descriptionKZ
	}

	return promotion.descriptionRU
}

func (promotion Promotion) Model() models.PromotionGet {
	return models.PromotionGet{
		ID:            promotion.ID(),
		Code:          promotion.Code(),
		TitleRU:       promotion.Title(languages.RU),
		TitleKZ:       promotion.Title(languages.KZ),
		ImageRU:       promotion.Image(languages.RU),
		DescriptionRU: promotion.Description(languages.RU),
		DescriptionKZ: promotion.Description(languages.KZ),
	}
}
