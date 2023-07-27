package brand_adaptor

import (
	"wkey-stock/src/data/dtos"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

func EntityToDTO(entityList []entities.Brand) []dtos.Brand {
	dtoList := make([]dtos.Brand, 0, len(entityList))

	for _, brand := range entityList {
		dtoList = append(dtoList, *dtos.NewBrand(&brand))
	}

	return dtoList
}

func DtoToModel(dtoList []dtos.Brand) []models.Brand {
	modelList := make([]models.Brand, 0, len(dtoList))

	for _, brand := range dtoList {
		modelList = append(modelList, brand.Model())
	}

	return modelList
}
