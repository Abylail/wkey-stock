package promotion_adaptor

import (
	"wkey-stock/src/data/dtos"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

func EntityToDTO(entityList []entities.AdminPromotion) []dtos.Promotion {
	dtoList := make([]dtos.Promotion, 0, len(entityList))

	for _, promotion := range entityList {
		dtoList = append(dtoList, dtos.NewPromotion(promotion))
	}

	return dtoList
}

func DtoToModel(dtoList []dtos.Promotion) []models.PromotionGet {
	modelList := make([]models.PromotionGet, 0, len(dtoList))

	for _, promotion := range dtoList {
		modelList = append(modelList, models.NewPromotion(promotion))
	}

	return modelList
}
