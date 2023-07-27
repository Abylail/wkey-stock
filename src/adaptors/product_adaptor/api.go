package product_adaptor

import (
	"wkey-stock/src/data/dtos"
	"wkey-stock/src/data/entities"
)

func EntityToDTO(entityList []entities.AdminProductGet) []dtos.Product {
	dtoList := make([]dtos.Product, 0, len(entityList))

	for _, product := range entityList {
		dtoList = append(dtoList, *dtos.NewProduct(&product))
	}

	return dtoList
}
