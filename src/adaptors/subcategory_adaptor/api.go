package subcategory_adaptor

import (
	"wkey-stock/src/data/dtos"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

func EntityToDTO(entityList []entities.SubCategory) []dtos.SubCategory {
	dtoList := make([]dtos.SubCategory, 0, len(entityList))

	for _, subCategory := range entityList {
		dtoList = append(dtoList, *dtos.NewSubCategory(&subCategory))
	}

	return dtoList
}

func DtoToModel(dtoList []dtos.SubCategory) []models.SubCategory {
	modelList := make([]models.SubCategory, 0, len(dtoList))

	for _, subCategory := range dtoList {
		modelList = append(modelList, *subCategory.Model())
	}

	return modelList
}
