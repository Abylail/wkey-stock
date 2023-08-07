package promotion_adaptor

import (
	"wkey-stock/src/data/dtos"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

func EntityToDTO(entityList []entities.Promotion) []dtos.Promotion {
	dtoList := make([]dtos.Promotion, 0, len(entityList))

	for _, promotion := range entityList {
		dtoList = append(dtoList, *dtos.NewPromotion(&promotion))
	}

	return dtoList
}

func DtoToModel(dtoList []dtos.Promotion) []models.Promotion {
	modelList := make([]models.Promotion, 0, len(dtoList))

	for _, promotion := range dtoList {
		modelList = append(modelList, *promotion.Model())
	}

	return modelList
}
