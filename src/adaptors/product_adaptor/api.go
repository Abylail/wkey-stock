package product_adaptor

import (
	"wkey-stock/src/data/dtos"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

func EntityToDTO(entityList []entities.Product) []dtos.Product {
	dtoList := make([]dtos.Product, 0, len(entityList))

	for _, product := range entityList {
		dtoList = append(dtoList, *dtos.NewProduct(&product))
	}

	return dtoList
}

func DtoToModel(dtoList []dtos.Product) []models.Product {
	modelList := make([]models.Product, 0, len(dtoList))

	for _, product := range dtoList {
		modelList = append(modelList, product.Model())
	}

	return modelList
}
