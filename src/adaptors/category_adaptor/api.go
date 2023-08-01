package category_adaptor

import (
	"wkey-stock/src/data/dtos"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

func EntityToDTO(entityList []entities.Category) []dtos.Category {
	dtoList := make([]dtos.Category, 0, len(entityList))

	for _, product := range entityList {
		dtoList = append(dtoList, *dtos.NewCategory(&product))
	}

	return dtoList
}

func DtoToModel(dtoList []dtos.Category) []models.Category {
	modelList := make([]models.Category, 0, len(dtoList))

	for _, product := range dtoList {
		modelList = append(modelList, *product.Model())
	}

	return modelList
}
